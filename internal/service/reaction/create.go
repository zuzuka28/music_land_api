package reaction

import (
	"context"
	"fmt"

	"github.com/zuzuka28/music_land_api/internal/model"
)

type creator interface {
	Create(ctx context.Context, item *model.Reaction) error
}

type createService struct {
	r creator
}

func newCreateService(r creator) *createService {
	return &createService{
		r: r,
	}
}

func (s *createService) Create(ctx context.Context, cmd *model.ReactionCreateCommand) error {
	if err := s.r.Create(ctx, &model.Reaction{
		UserID:             cmd.UserID,
		ReactionType:       cmd.ReactionType,
		ReactionTargetID:   cmd.ReactionTargetID,
		ReactionTargetType: cmd.ReactionTargetType,
	}); err != nil {
		return fmt.Errorf("create reaction: %w", err)
	}

	return nil
}
