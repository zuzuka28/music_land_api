package auth

import (
	"context"

	"github.com/zuzuka28/music_land_api/internal/model"
	"go.opentelemetry.io/otel/trace"
)

type (
	Tracer interface {
		Start(ctx context.Context, spanName string) (context.Context, trace.Span)
	}

	UserService interface {
		Fetch(ctx context.Context, query *model.UserFetchQuery) (*model.User, error)
	}
)
