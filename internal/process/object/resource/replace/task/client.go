package task

import (
	"context"

	"github.com/spacelift-io/homework-object-storage/internal/process/object/resource/replace/task/client"
	"github.com/spacelift-io/homework-object-storage/internal/service/storage"
)

func Client(c context.Context, input client.Input) (client.Output, error) {
	return client.Output{
		Client: storage.ClientByID(input.ID),
	}, nil
}
