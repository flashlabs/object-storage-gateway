package task

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"

	"github.com/spacelift-io/homework-object-storage/internal/process/object/resource/read/task/checkobject"
	"github.com/spacelift-io/homework-object-storage/internal/service/storage"
)

var (
	objNotExistsErrorMessage = "The specified key does not exist."
)

func CheckObject(c context.Context, input checkobject.Input) (checkobject.Output, error) {
	_, err := input.Client.StatObject(c, storage.BucketName, input.ID, minio.StatObjectOptions{})
	if err != nil && err.Error() != objNotExistsErrorMessage {
		return checkobject.Output{}, fmt.Errorf("error while executing input.Client.StatObject: %w", err)
	}

	ex := !(err != nil && err.Error() == objNotExistsErrorMessage)

	return checkobject.Output{
		Exists: ex,
	}, nil
}
