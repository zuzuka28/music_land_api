package track

import (
	"context"

	"github.com/zuzuka28/music_land_api/internal/model"
	"github.com/zuzuka28/music_land_api/pkg/fs"
)

type (
	Service interface {
		Fetch(ctx context.Context, query *model.TrackFetchQuery) (*model.Track, error)
		FetchFile(ctx context.Context, query *model.TrackFetchQuery) (*fs.File, error)
		Search(ctx context.Context, query *model.TrackSearchQuery) ([]*model.Track, error)
		Create(ctx context.Context, cmd *model.TrackCreateCommand) error
		Delete(ctx context.Context, cmd *model.TrackDeleteCommand) error
	}
)
