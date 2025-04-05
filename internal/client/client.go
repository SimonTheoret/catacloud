package client

import (
	"fmt"
	"time"
)

// This struct contains the data we need from files located on the cloud.
type FileWithMetadata struct {
	LastModified time.Time
	Name         string
}

// Output of a `ListFiles` call from a `Client`
type ListedFiles[T fmt.Stringer] struct {
	Inner []T
	Files []FileWithMetadata
}

// Client interface over what a client must do and return.
type Client[T fmt.Stringer] interface {
	ListFiles() ListedFiles[fmt.Stringer]
}

// Client interface for downloading files
type Downlader interface {
	DownloadFiles(files []string) (err error)
}

// Client interface for uploading files
type Uploader interface {
	UploadFiles(filepaths []string) (err error)
}
