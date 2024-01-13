package initializer

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"github.com/spacelift-io/homework-object-storage/internal/registry"
	"github.com/spacelift-io/homework-object-storage/internal/structs"
	"github.com/spacelift-io/homework-object-storage/pkg"
)

var (
	containerNumRegex = regexp.MustCompile(`[0-9]+`)
	image             = "minio/minio"
)

func Shards(c context.Context) error {
	log.Println("Initializing shards")

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return fmt.Errorf("error while executing client.NewClientWithOpts: %w", err)
	}

	// TODO: filter by image
	containers, err := cli.ContainerList(c, types.ContainerListOptions{})
	if err != nil {
		return fmt.Errorf("error while executing cli.ContainerList: %w", err)
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

		shards[uint8(containerNum)] = structs.Storage{
			Container: container,
		}
	}

	if len(shards) == 0 {
		return pkg.ErrMissingStorageContainers
	}

	registry.Shards = shards

	return nil
}
