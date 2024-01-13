package structs

import (
	"github.com/minio/minio-go/v7"
)

type Storage struct {
	Client *minio.Client
}
