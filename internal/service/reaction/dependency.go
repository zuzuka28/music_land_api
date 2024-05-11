package reaction

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
		Create(ctx context.Context, item *model.Reaction) error
		Search(ctx context.Context, query *model.ReactionSearchQuery) ([]*model.Reaction, error)
		Delete(ctx context.Context, cmd *model.ReactionDeleteCommand) error
	}
)
