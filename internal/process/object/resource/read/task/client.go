package task

import (
	"context"
	"fmt"

	"github.com/spacelift-io/homework-object-storage/internal/process/object/resource/read/task/client"
	"github.com/spacelift-io/homework-object-storage/internal/service/storage"
)

func Client(c context.Context, input client.Input) (client.Output, error) {
	sc, err := storage.ClientByID(input.ID)
	if err != nil {
		return client.Output{}, fmt.Errorf("error while executing storage.ClientByID: %w", err)
	}

	return client.Output{
		Client: sc,
	}, nil
}
