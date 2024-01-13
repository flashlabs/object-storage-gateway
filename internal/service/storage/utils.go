package storage

import (
	"hash/fnv"

	"github.com/minio/minio-go/v7"

	"github.com/spacelift-io/homework-object-storage/internal/registry"
)

const (
	BucketName    = "default"
	RegionName    = "eu-central-1"
	ObjectLocking = true
)

func ClientByID(id string) *minio.Client {
	return registry.Shards[ShardByID(id)].Client
}

func ShardByID(id string) uint8 {
	hash := fnv.New32a()
	hash.Write([]byte(id))
	checksum := hash.Sum32()
	shardIndex := checksum % uint32(len(registry.Shards))

	return uint8(shardIndex)
}
