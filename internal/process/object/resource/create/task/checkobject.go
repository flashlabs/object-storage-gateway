package task

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"

	"github.com/spacelift-io/homework-object-storage/internal/process/object/resource/create/task/checkobject"
	"github.com/spacelift-io/homework-object-storage/internal/service/storage"
)

func CheckObject(c context.Context, input checkobject.Input) (checkobject.Output, error) {
	o, err := input.Client.GetObject(c, storage.BucketName, input.ID, minio.GetObjectOptions{})
	if err != nil {
		return checkobject.Output{}, fmt.Errorf("error while executing input.Client.GetObject: %w", err)
	}

	defer func(o *minio.Object) {
		if ce := o.Close(); ce != nil {
			err = ce
		}
	}(o)

	return checkobject.Output{
		Exists: true,
	}, err
}
