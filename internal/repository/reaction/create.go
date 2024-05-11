package reaction

import (
	"context"
	"fmt"

	"github.com/zuzuka28/music_land_api/internal/model"
	"xorm.io/xorm"
)

type createRepo struct {
	db *xorm.Engine
}

func newCreateRepository(db *xorm.Engine) *createRepo {
	return &createRepo{
		db: db,
	}
}

func (r *createRepo) Create(ctx context.Context, item *model.Reaction) error {
	react := mapReactionToInternal(item)

	_, err := r.db.Context(ctx).Insert(react)
	if err != nil {
		return fmt.Errorf("create reaction: %w", err)
	}

	return nil
}
