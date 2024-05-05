package fs

import (
	"io"
	"time"
)

type File struct {
	Name         string
	Size         int64
	LastModified time.Time
	Reader       io.Reader
}
