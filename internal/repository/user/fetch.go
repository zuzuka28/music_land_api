package user

import (
	"context"
	"fmt"

	"github.com/zuzuka28/music_land_api/internal/model"
	"xorm.io/xorm"
)

type fetchRepo struct {
	db *xorm.Engine
}

func newFetchRepository(db *xorm.Engine) *fetchRepo {
	return &fetchRepo{
		db: db,
	}
}

func (r *fetchRepo) Fetch(ctx context.Context, id string) (*model.User, error) {
	usr := &user{ //nolint:exhaustruct
		Username: id,
	}

	has, err := r.db.Context(ctx).Get(usr)
	if err != nil {
		return nil, fmt.Errorf("get user from db: %w", err)
	}

	if !has {
		return nil, fmt.Errorf("%w: user %s", model.ErrNotFound, id)
	}

	return mapUserToModel(usr), nil
}
