package client

import "github.com/minio/minio-go/v7"

type Output struct {
	Client *minio.Client
}
