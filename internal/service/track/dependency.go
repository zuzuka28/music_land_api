package track

import (
	"context"

	"github.com/zuzuka28/music_land_api/internal/model"
)

type (
	FileStorage interface {
		SaveFile(ctx context.Context, name string, data []byte) error
		DeleteFile(ctx context.Context, name string) error
	}

	Repository interface {
		Create(ctx context.Context, item *model.Track) error
		Fetch(ctx context.Context, id string) (*model.Track, error)
		Delete(ctx context.Context, id string) error
		Search(ctx context.Context, query *model.TrackSearchQuery) ([]*model.Track, error)
	}
)
