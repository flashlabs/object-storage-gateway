package task

import (
	"context"
	"fmt"

	"github.com/spacelift-io/homework-object-storage/internal/process/object/resource/replace/task/checkbucket"
	"github.com/spacelift-io/homework-object-storage/internal/service/storage"
)

func CheckBucket(c context.Context, input checkbucket.Input) (checkbucket.Output, error) {
	found, err := input.Client.BucketExists(c, storage.BucketName)
	if err != nil {
		return checkbucket.Output{}, fmt.Errorf("error while executing input.Client.BucketExists: %w", err)
	}

	return checkbucket.Output{
		Exists: found,
	}, nil
}
