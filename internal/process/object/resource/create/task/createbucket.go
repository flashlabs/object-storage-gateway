package task

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"

	"github.com/spacelift-io/homework-object-storage/internal/process/object/resource/create/task/createbucket"
	"github.com/spacelift-io/homework-object-storage/internal/service/storage"
)

func CreateBucket(c context.Context, input createbucket.Input) (createbucket.Output, error) {
	err := input.Client.MakeBucket(c, storage.BucketName, minio.MakeBucketOptions{
		Region:        storage.RegionName,
		ObjectLocking: storage.ObjectLocking,
	})
	if err != nil {
		return createbucket.Output{}, fmt.Errorf("error while executing input.Client.MakeBucket: %w", err)
	}

	return createbucket.Output{}, nil
}
