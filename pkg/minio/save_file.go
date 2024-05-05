package minio

import (
	"bytes"
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

type saveClient struct {
	s3     *minio.Client
	bucket string
}

func newSaveClient(s3 *minio.Client, bucket string) *saveClient {
	return &saveClient{
		s3:     s3,
		bucket: bucket,
	}
}

func (c *saveClient) SaveFile(ctx context.Context, name string, data []byte) error {
	r := bytes.NewReader(data)

	_, err := c.s3.PutObject(
		ctx,
		c.bucket,
		name,
		r,
		r.Size(),
		minio.PutObjectOptions{}, //nolint:exhaustruct
	)
	if err != nil {
		return fmt.Errorf("put file to storage: %w", err)
	}

	return nil
}
