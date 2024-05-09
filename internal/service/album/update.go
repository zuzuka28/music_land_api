package album

import (
	"context"
	"fmt"

	"github.com/zuzuka28/music_land_api/internal/model"
)

type fetcherUpdater interface {
	Fetch(ctx context.Context, id string) (*model.Album, error)
	Update(ctx context.Context, cmd *model.AlbumUpdateCommand) error
}

type updateService struct {
	r fetcherUpdater
}

func newUpdateService(r fetcherUpdater) *updateService {
	return &updateService{
		r: r,
	}
}

func (s *updateService) Update(ctx context.Context, cmd *model.AlbumUpdateCommand) error {
	if err := s.r.Update(ctx, cmd); err != nil {
		return fmt.Errorf("update album: %w", err)
	}

	return nil
}
