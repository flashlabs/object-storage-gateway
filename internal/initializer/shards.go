package initializer

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var (
	reContainerNum = regexp.MustCompile(`[0-9]+`)
)

func Shards(c context.Context) error {
	log.Println("Initializing shards")

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return fmt.Errorf("error while executing client.NewClientWithOpts: %w", err)
	}

	// TODO: filter by name
	containers, err := cli.ContainerList(c, types.ContainerListOptions{})
	if err != nil {
		return fmt.Errorf("error while executing cli.ContainerList: %w", err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s %s\n", container.ID[:10], container.Image, container.Names)

		name := container.Names[0]

		fmt.Println("name2 ", reContainerNum.FindString(name))
	}

	return nil
}
