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
	_, err := r.db.Transaction(func(sess *xorm.Session) (interface{}, error) {
		sess = sess.Context(ctx)

		if _, err := sess.Delete(&reaction{ //nolint:exhaustruct
			UserID:           cmd.UserID,
			ReactionTargetID: cmd.ReactionTargetID,
		}); err != nil {
			return nil, fmt.Errorf("delete reaction in db: %w", err)
		}

		return nil, nil
	})
	if err != nil {
		return fmt.Errorf("delete reaction: %w", err)
	}

	return nil
}
