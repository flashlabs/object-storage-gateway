package object

import "github.com/minio/minio-go/v7"

type Input struct {
	Client *minio.Client
	ID     string
}
