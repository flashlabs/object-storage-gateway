package pkg

import "errors"

var (
	ErrMissingStorageContainers   = errors.New("no storage container was found, no shards were created")
	ErrContainerIPAddressNotFound = errors.New("IP address for container not found")
	ErrObjectNotExists            = errors.New("object not found")
	ErrNoStorageClientForGivenID  = errors.New("no storage client for given ID")
)
