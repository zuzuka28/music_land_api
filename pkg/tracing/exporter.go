package tracing

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
)

func NewOTLPExporter(ctx context.Context, cfg ExporterConfig) (*otlptrace.Exporter, error) {
	client := otlptracegrpc.NewClient(
		otlptracegrpc.WithEndpoint(cfg.URL),
		otlptracegrpc.WithInsecure(),
	)

	exp, err := otlptrace.New(ctx, client)
	if err != nil {
		return nil, fmt.Errorf("new otlptrace: %w", err)
	}

	return exp, nil
}
