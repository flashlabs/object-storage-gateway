package pkg

import "errors"

var (
	ErrMissingStorageContainers = errors.New("no storage container was found, no shards were created")
)
