package album

import (
	"context"
	"fmt"

	"github.com/zuzuka28/music_land_api/internal/model"
)

type creator interface {
	Create(ctx context.Context, item *model.Album) error
}

type createService struct {
	r creator
}

func newCreateService(r creator) *createService {
	return &createService{
		r: r,
	}
}

func (s *createService) Create(ctx context.Context, cmd *model.AlbumCreateCommand) error {
	if err := s.r.Create(ctx, &model.Album{
		ID:       "",
		Name:     cmd.Name,
		OwnerID:  cmd.OwnerID,
		TrackIDs: cmd.TrackIDs,
	}); err != nil {
		return fmt.Errorf("create album: %w", err)
	}

	return nil
}
