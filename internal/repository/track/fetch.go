package track

import (
	"context"
	"fmt"
	"strconv"

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

func (r *fetchRepo) Fetch(ctx context.Context, id string) (*model.Track, error) {
	tid, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", model.ErrNotValid, err)
	}

	tr := &track{ //nolint:exhaustruct
		ID: int64(tid),
	}

	has, err := r.db.Context(ctx).Get(tr)
	if err != nil {
		return nil, fmt.Errorf("get track from db: %w", err)
	}

	if !has {
		return nil, fmt.Errorf("%w: track %s", model.ErrNotFound, id)
	}

	return mapTrackToModel(tr), nil
}
