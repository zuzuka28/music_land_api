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

	Repository interface {
		Fetch(ctx context.Context, id string) (*model.User, error)
		Create(ctx context.Context, item *model.User) error
	}
)
