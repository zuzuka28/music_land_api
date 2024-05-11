package album

import (
	"context"
	"fmt"

	"github.com/google/uuid"
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

func (r *createRepo) Create(ctx context.Context, item *model.Album) error {
	if item.ID == "" {
		item.ID = uuid.New().String()
	}

	alb := mapAlbumToInternal(item)
	items := mapAlbumItems(item)

	_, err := r.db.Transaction(func(sess *xorm.Session) (any, error) {
		sess = sess.Context(ctx)

		if _, err := sess.Insert(alb); err != nil {
			return nil, fmt.Errorf("insert album in db: %w", err)
		}

		if _, err := sess.Insert(items); err != nil {
			return nil, fmt.Errorf("insert album items in db: %w", err)
		}

		return nil, nil //nolint:nilnil
	})
	if err != nil {
		return fmt.Errorf("create album: %w", err)
	}

	return nil
}
