package minio

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

type deleteClient struct {
	s3 *minio.Client

	bucket string
}

func newDeleteClient(s3 *minio.Client, bucket string) *deleteClient {
	return &deleteClient{
		s3:     s3,
		bucket: bucket,
	}
}

func (c *deleteClient) DeleteFile(ctx context.Context, name string) error {
	if err := c.s3.RemoveObject(
		ctx,
		c.bucket,
		name,
		minio.RemoveObjectOptions{}, //nolint:exhaustruct
	); err != nil {
		return fmt.Errorf("remove object from storage: %w", err)
	}

	return nil
}
