package minio

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/zuzuka28/music_land_api/pkg/fs"
)

type Credentials struct {
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
}

type Config struct {
	Endpoint    string       `yaml:"endpoint"`
	Bucket      string       `yaml:"bucket"`
	Credentials *Credentials `yaml:"credentials"`
}

type Client struct {
	tr Tracer

	dc *deleteClient
	sc *saveClient
	fc *fetchClient
}

func NewClient(cfg *Config, tr Tracer) (*Client, error) {
	s3, err := minio.New(cfg.Endpoint, &minio.Options{ //nolint:exhaustruct
		Creds: credentials.NewStaticV4(
			cfg.Credentials.AccessKey,
			cfg.Credentials.SecretKey,
			"",
		),
	})
	if err != nil {
		return nil, fmt.Errorf("create minio client: %w", err)
	}

	return &Client{
		tr: tr,
		dc: newDeleteClient(s3, cfg.Bucket),
		sc: newSaveClient(s3, cfg.Bucket),
		fc: newFetchClient(s3, cfg.Bucket),
	}, nil
}

func (c *Client) DeleteFile(ctx context.Context, name string) error {
	ctx, span := c.tr.Start(ctx, "DeleteFile")
	defer span.End()

	return c.dc.DeleteFile(ctx, name)
}

func (c *Client) FetchFile(ctx context.Context, name string) (*fs.File, error) {
	ctx, span := c.tr.Start(ctx, "FetchFile")
	defer span.End()

	return c.fc.FetchFile(ctx, name)
}

func (c *Client) SaveFile(ctx context.Context, name string, data []byte) error {
	ctx, span := c.tr.Start(ctx, "SaveFile")
	defer span.End()

	return c.sc.SaveFile(ctx, name, data)
}
