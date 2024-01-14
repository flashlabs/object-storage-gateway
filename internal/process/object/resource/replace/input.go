package replace

import "io"

type Input struct {
	Vars          map[string]string
	Payload       io.Reader
	ContentLength int64
}
