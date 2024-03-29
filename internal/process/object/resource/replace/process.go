package replace

import (
	"context"
	"fmt"

	"github.com/spacelift-io/homework-object-storage/internal/process/object/resource/replace/task"
	"github.com/spacelift-io/homework-object-storage/internal/process/object/resource/replace/task/checkbucket"
	"github.com/spacelift-io/homework-object-storage/internal/process/object/resource/replace/task/checkobject"
	"github.com/spacelift-io/homework-object-storage/internal/process/object/resource/replace/task/client"
	"github.com/spacelift-io/homework-object-storage/internal/process/object/resource/replace/task/createbucket"
	"github.com/spacelift-io/homework-object-storage/internal/process/object/resource/replace/task/upload"
)

func Execute(c context.Context, input Input) (Output, error) {
	// get storage client
	clientOut, err := task.Client(client.Input{ID: input.Vars["id"]})
	if err != nil {
		return Output{}, fmt.Errorf("error while executing task.Client: %w", err)
	}

	mc := clientOut.Client

	checkBucketOut, err := task.CheckBucket(c, checkbucket.Input{Client: mc})
	if err != nil {
		return Output{}, fmt.Errorf("error while executing task.CheckBucket: %w", err)
	}

	if !checkBucketOut.Exists {
		_, err := task.CreateBucket(c, createbucket.Input{Client: mc})
		if err != nil {
			return Output{}, fmt.Errorf("error while executing task.CreateBucket: %w", err)
		}
	}

	checkObjOut, err := task.CheckObject(c, checkobject.Input{
		Client: mc,
		ID:     input.Vars["id"],
	})
	if err != nil {
		return Output{}, fmt.Errorf("error while executing task.CheckObject: %w", err)
	}

	_, err = task.Upload(c, upload.Input{
		Client:        mc,
		ID:            input.Vars["id"],
		Payload:       input.Payload,
		ContentLength: input.ContentLength,
	})
	if err != nil {
		return Output{}, fmt.Errorf("error while executing task.Upload: %w", err)
	}

	return Output{
		Replaced: checkObjOut.Exists,
	}, nil
}
