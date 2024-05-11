package reaction

import (
	"context"
	"fmt"

	"github.com/zuzuka28/music_land_api/internal/model"
	"xorm.io/xorm"
)

type deleteRepo struct {
	db *xorm.Engine
}

func newDeleteRepository(db *xorm.Engine) *deleteRepo {
	return &deleteRepo{
		db: db,
	}
}

func (r *deleteRepo) Delete(ctx context.Context, cmd *model.ReactionDeleteCommand) error {
	react := &reaction{ //nolint:exhaustruct
		UserID:           cmd.UserID,
		ReactionTargetID: cmd.ReactionTargetID,
	}

	_, err := r.db.Context(ctx).Delete(react)
	if err != nil {
		return fmt.Errorf("delete reaction: %w", err)
	}

	return nil
}
