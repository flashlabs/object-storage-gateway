package task

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"

	"github.com/spacelift-io/homework-object-storage/internal/process/object/resource/replace/task/upload"
	"github.com/spacelift-io/homework-object-storage/internal/service/storage"
)

const (
	contentType = "application/octet-stream"
)

func Upload(c context.Context, input upload.Input) (upload.Output, error) {
	_, err := input.Client.PutObject(c, storage.BucketName, input.ID, input.Payload, input.ContentLength, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return upload.Output{}, fmt.Errorf("error while executing input.Client.PutObject: %w", err)
	}

	return upload.Output{}, nil
}
