package user

import (
	"context"

	"github.com/zuzuka28/music_land_api/internal/model"
	"go.opentelemetry.io/otel/trace"
)

type (
	Tracer interface {
		Start(ctx context.Context, spanName string) (context.Context, trace.Span)
	}

	Service interface {
		Create(ctx context.Context, cmd *model.UserCreateCommand) error
	}
)
