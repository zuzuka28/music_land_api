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

	Service interface {
		Create(ctx context.Context, cmd *model.AlbumCreateCommand) error
		Delete(ctx context.Context, cmd *model.AlbumDeleteCommand) error
		Fetch(ctx context.Context, query *model.AlbumFetchQuery) (*model.Album, error)
		Update(ctx context.Context, cmd *model.AlbumUpdateCommand) error
	}
)
