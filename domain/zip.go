package domain

import "io"

type IZipService interface {
	Unzip([]byte, string) ([]File, error)
}

// File stores a files' name and body.
type File struct {
	Name string
	Body io.ReadCloser
}
