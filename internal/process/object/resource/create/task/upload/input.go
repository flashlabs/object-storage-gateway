package upload

import (
	"io"

	"github.com/minio/minio-go/v7"
)

type Input struct {
	Payload io.Reader
	Client  *minio.Client
	ID      string
}
