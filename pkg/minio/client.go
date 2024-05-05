package minio

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
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
	*deleteClient
	*saveClient
}

func NewClient(cfg *Config) (*Client, error) {
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
		deleteClient: newDeleteClient(s3, cfg.Bucket),
		saveClient:   newSaveClient(s3, cfg.Bucket),
	}, nil
}

func (c *Client) DeleteFile(ctx context.Context, name string) error {
	return c.deleteClient.DeleteFile(ctx, name)
}

func (c *Client) SaveFile(ctx context.Context, name string, data []byte) error {
	return c.saveClient.SaveFile(ctx, name, data)
}
