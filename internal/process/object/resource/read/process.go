package read

import (
	"context"
	"fmt"

	"github.com/spacelift-io/homework-object-storage/internal/process/object/resource/read/task"
	"github.com/spacelift-io/homework-object-storage/internal/process/object/resource/read/task/checkobject"
	"github.com/spacelift-io/homework-object-storage/internal/process/object/resource/read/task/client"
	"github.com/spacelift-io/homework-object-storage/internal/process/object/resource/read/task/object"
	"github.com/spacelift-io/homework-object-storage/pkg"
)

func Execute(c context.Context, input Input) (Output, error) {
	// get storage client
	clientOut, err := task.Client(client.Input{ID: input.Vars["id"]})
	if err != nil {
		return Output{}, fmt.Errorf("error while executing task.Client: %w", err)
	}

	mc := clientOut.Client

	checkObjOut, err := task.CheckObject(c, checkobject.Input{
		Client: mc,
		ID:     input.Vars["id"],
	})
	if err != nil {
		return Output{}, fmt.Errorf("error while executing task.CheckObject: %w", err)
	}

	if !checkObjOut.Exists {
		return Output{}, pkg.ErrObjectNotExists
	}

	readOut, err := task.Object(c, object.Input{
		ID:     input.Vars["id"],
		Client: mc,
	})
	if err != nil {
		return Output{}, fmt.Errorf("error while executing task.Object: %w", err)
	}

	return Output{
		Entity: readOut.Entity,
	}, nil
}
