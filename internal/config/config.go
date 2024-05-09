package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/zuzuka28/music_land_api/pkg/minio"
	"github.com/zuzuka28/music_land_api/pkg/tracing"
)

type Service struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Storage struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type API struct {
	Service     *Service        `yaml:"service"`
	FileStorage *minio.Config   `yaml:"file_storage"`
	Storage     *Storage        `yaml:"storage"`
	Tracing     *tracing.Config `yaml:"tracing"`
	LogLevel    string          `yaml:"log_level"`
}

func NewAPI(configPath string) (*API, error) {
	cfg := new(API)

	err := cleanenv.ReadConfig(configPath, cfg)
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	return cfg, nil
}
