package minio

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/zuzuka28/music_land_api/pkg/fs"
)

type fetchClient struct {
	s3     *minio.Client
	bucket string
}

func newFetchClient(s3 *minio.Client, bucket string) *fetchClient {
	return &fetchClient{
		s3:     s3,
		bucket: bucket,
	}
}

func (c *fetchClient) FetchFile(ctx context.Context, name string) (*fs.File, error) {
	info, err := c.s3.GetObjectAttributes(
		ctx,
		c.bucket,
		name, minio.ObjectAttributesOptions{}, //nolint:exhaustruct
	)
	if err != nil {
		return nil, fmt.Errorf("get file info from storage: %w", err)
	}

	res, err := c.s3.GetObject(
		ctx, c.bucket, name, minio.GetObjectOptions{}) //nolint:exhaustruct
	if err != nil {
		return nil, fmt.Errorf("get file from storage: %w", err)
	}

	return &fs.File{
		Name:         name,
		Size:         int64(info.ObjectSize),
		LastModified: info.LastModified,
		Reader:       res,
	}, nil
}
