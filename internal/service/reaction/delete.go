package reaction

import (
	"context"
	"fmt"

	"github.com/zuzuka28/music_land_api/internal/model"
)

type deleter interface {
	Delete(ctx context.Context, cmd *model.ReactionDeleteCommand) error
}

type deleteService struct {
	r deleter
}

func newDeleteService(r deleter) *deleteService {
	return &deleteService{
		r: r,
	}
}

func (s *deleteService) Delete(ctx context.Context, cmd *model.ReactionDeleteCommand) error {
	if err := s.r.Delete(ctx, cmd); err != nil {
		return fmt.Errorf("delete reaction: %w", err)
	}

	return nil
}
