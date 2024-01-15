package initializer

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/spacelift-io/homework-object-storage/internal/registry"
	"github.com/spacelift-io/homework-object-storage/internal/structs"
	"github.com/spacelift-io/homework-object-storage/pkg"
)

var (
	containerNumRegex = regexp.MustCompile(`[0-9]+`)
)

const (
	apiPort       = 9000
	containerName = "amazin-object-storage-node"
	image         = "minio/minio"
)

type connParams struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	SSL       bool
}

func Shards(c context.Context) error {
	log.Println("Initializing shards")

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return fmt.Errorf("error while executing client.NewClientWithOpts: %w", err)
	}

	containers, err := listContainers(c, cli)
	if err != nil {
		return fmt.Errorf("error while executing listContainers: %w", err)
	}

	shards := make(map[uint8]structs.Storage)

	for _, container := range containers {
		if container.Image != image {
			continue
		}

		numFound := containerNumRegex.FindString(container.Names[0])
		if numFound == "" {
			continue
		}

		containerNum, err := strconv.ParseInt(numFound, 10, 8)
		if err != nil {
			return fmt.Errorf("error while executing strconv.ParseInt: %w", err)
		}

		log.Println("Found container", container.ID[:12])

		cp, err := connectionParams(c, cli, container)
		if err != nil {
			return fmt.Errorf("error while executing connectionParams: %w", err)
		}

		minioClient, err := minio.New(cp.Endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(cp.AccessKey, cp.SecretKey, ""),
			Secure: cp.SSL,
		})
		if err != nil {
			return fmt.Errorf("error while executing minio.New: %w", err)
		}

		// shards indexes are zero based
		shardID := uint8(containerNum - 1)
		shards[shardID] = structs.Storage{
			Client: minioClient,
		}
	}

	if len(shards) == 0 {
		return pkg.ErrMissingStorageContainers
	}

	registry.Shards = shards

	return nil
}

func listContainers(c context.Context, cli *client.Client) ([]types.Container, error) {
	args := filters.NewArgs()
	args.Add("name", containerName)

	list, err := cli.ContainerList(c, types.ContainerListOptions{
		All:     true, // Include stopped containers as well
		Filters: args},
	)
	if err != nil {
		return nil, fmt.Errorf("error while executing cli.ContainerList: %w", err)
	}

	return list, nil
}

func connectionParams(c context.Context, cli *client.Client, container types.Container) (connParams, error) {
	info, err := cli.ContainerInspect(c, container.ID)
	if err != nil {
		return connParams{}, fmt.Errorf("error while executing cli.ContainerInspect: %w", err)
	}

	envs := make(map[string]string)
	containerEnvVars(info, envs)

	address := containerIPAddress(info)
	if address == "" {
		return connParams{}, pkg.ErrContainerIPAddressNotFound
	}

	return connParams{
		Endpoint:  fmt.Sprintf("%s:%d", address, apiPort),
		AccessKey: envs["MINIO_ACCESS_KEY"],
		SecretKey: envs["MINIO_SECRET_KEY"],
		SSL:       false,
	}, nil
}

func containerEnvVars(info types.ContainerJSON, envs map[string]string) {
	for _, v := range info.Config.Env {
		parts := strings.SplitN(v, "=", 2)
		if len(parts) == 2 {
			envs[parts[0]] = parts[1]
		}
	}
}

func containerIPAddress(info types.ContainerJSON) string {
	for _, nw := range info.NetworkSettings.Networks {
		return nw.IPAddress
	}

	return ""
}
