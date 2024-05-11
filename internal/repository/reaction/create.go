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

	_, err := r.db.Transaction(func(sess *xorm.Session) (interface{}, error) {
		sess = sess.Context(ctx)

		if _, err := sess.Insert(react); err != nil {
			return nil, fmt.Errorf("insert reaction in db: %w", err)
		}

		return nil, nil
	})
	if err != nil {
		return fmt.Errorf("create reaction: %w", err)
	}

	return nil
}
