package album

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
		Create(ctx context.Context, item *model.Album) error
		Fetch(ctx context.Context, id string) (*model.Album, error)
		Delete(ctx context.Context, id string) error
		Update(ctx context.Context, cmd *model.AlbumUpdateCommand) error
	}
)
