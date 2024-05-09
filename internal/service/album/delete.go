package album

import (
	"context"
	"fmt"

	"github.com/zuzuka28/music_land_api/internal/model"
)

type deleter interface {
	Delete(ctx context.Context, id string) error
}

type deleteService struct {
	r deleter
}

func newDeleteService(r Repository) *deleteService {
	return &deleteService{
		r: r,
	}
}

func (s *deleteService) Delete(ctx context.Context, cmd *model.AlbumDeleteCommand) error {
	if err := s.r.Delete(ctx, cmd.AlbumID); err != nil {
		return fmt.Errorf("delete album: %w", err)
	}

	return nil
}
