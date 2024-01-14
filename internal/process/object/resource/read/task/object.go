package task

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"

	"github.com/spacelift-io/homework-object-storage/internal/process/object/resource/read/task/object"
	"github.com/spacelift-io/homework-object-storage/internal/service/storage"
)

func Object(c context.Context, input object.Input) (object.Output, error) {
	o, err := input.Client.GetObject(c, storage.BucketName, input.ID, minio.StatObjectOptions{})
	if err != nil {
		return object.Output{}, fmt.Errorf("error while executing input.Client.Object: %w", err)
	}

	return object.Output{
		Entity: o,
	}, nil
}
