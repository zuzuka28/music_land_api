package track

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

func (r *createRepo) Create(ctx context.Context, item *model.Track) error {
	usr := mapTrackToInternal(item)

	_, err := r.db.Context(ctx).Insert(usr)
	if err != nil {
		return fmt.Errorf("insert track to db: %w", err)
	}

	return nil
}
