package tracing

import (
	"context"

	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

const SpanNameSeparator = ":"

type Config struct {
	Exporter ExporterConfig `yaml:"exporter"`
	Resource ResourceConfig `yaml:"resource"`
}

type ExporterConfig struct {
	URL string `yaml:"url"`
}

type ResourceConfig struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version" default:"0.0.1"`
}

type Tracer struct {
	tracer   trace.Tracer
	provider trace.TracerProvider
	name     string
	cleanup  func(ctx context.Context) error
}

func New(exp sdktrace.SpanExporter, r *resource.Resource) *Tracer {
	p := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	)

	return &Tracer{
		tracer:   p.Tracer(""),
		provider: p,
		name:     "",
		cleanup:  exp.Shutdown,
	}
}

func NewNoop() *Tracer {
	p := noop.NewTracerProvider()

	return &Tracer{
		tracer:   p.Tracer("noop"),
		provider: p,
		name:     "noop",
		cleanup:  nil,
	}
}

func (tr *Tracer) Child(name string) *Tracer {
	childName := name

	if tr.name != "" {
		childName = tr.name + SpanNameSeparator + name
	}

	return &Tracer{
		tracer:   tr.Provider().Tracer(childName),
		provider: tr.Provider(),
		name:     childName,
		cleanup:  nil,
	}
}

func (tr *Tracer) Provider() trace.TracerProvider {
	return tr.provider
}

func (tr *Tracer) Start(ctx context.Context, spanName string) (context.Context, trace.Span) {
	return tr.tracer.Start(ctx, spanName) //nolint:spancheck
}

func (tr *Tracer) Shutdown(ctx context.Context) error {
	if tr.cleanup == nil {
		return nil
	}

	return tr.cleanup(ctx)
}

func NewResource(cfg ResourceConfig) *resource.Resource {
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(cfg.Name),
		semconv.ServiceVersionKey.String(cfg.Version),
	)
}
