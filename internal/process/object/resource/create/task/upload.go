package task

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"

	"github.com/spacelift-io/homework-object-storage/internal/process/object/resource/create/task/upload"
	"github.com/spacelift-io/homework-object-storage/internal/service/storage"
)

func Upload(c context.Context, input upload.Input) (upload.Output, error) {
	_, err := input.Client.PutObject(c, storage.BucketName, input.ID, input.Payload, payloadSize(input.Payload), minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	if err != nil {
		return upload.Output{}, fmt.Errorf("error while executing input.Client.PutObject: %w", err)
	}

	return upload.Output{}, nil
}

func payloadSize(stream io.Reader) int64 {
	buf := new(bytes.Buffer)

	_, err := buf.ReadFrom(stream)
	if err != nil {
		return 0
	}

	return int64(buf.Len())
}
