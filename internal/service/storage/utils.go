package storage

import (
	"fmt"
	"hash/fnv"

	"github.com/minio/minio-go/v7"

	"github.com/spacelift-io/homework-object-storage/internal/registry"
	"github.com/spacelift-io/homework-object-storage/pkg"
)

const (
	BucketName    = "default"
	RegionName    = "eu-central-1"
	ObjectLocking = true
)

func ClientByID(id string) (*minio.Client, error) {
	idx, err := ShardByID(id)
	if err != nil {
		return nil, fmt.Errorf("error while executing ShardByID: %w", err)
	}

	if s, ok := registry.Shards[idx]; ok {
		return s.Client, nil
	}

	return nil, pkg.ErrNoStorageClientForGivenID
}

func ShardByID(id string) (uint8, error) {
	if len(registry.Shards) == 0 {
		return 0, pkg.ErrNoShardsAvailable
	}

	hash := fnv.New32a()

	_, err := hash.Write([]byte(id))
	if err != nil {
		return 0, fmt.Errorf("error while executing hash.Write: %w", err)
	}

	checksum := hash.Sum32()
	shardIndex := checksum % uint32(len(registry.Shards))

	return uint8(shardIndex), nil
}
