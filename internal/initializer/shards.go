package initializer

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func Shards(c context.Context) error {
	log.Println("Initializing shards")

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return fmt.Errorf("error while executing client.NewClientWithOpts: %w", err)
	}

	containers, err := cli.ContainerList(c, types.ContainerListOptions{})
	if err != nil {
		return fmt.Errorf("error while executing cli.ContainerList: %w", err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}

	return nil
}
