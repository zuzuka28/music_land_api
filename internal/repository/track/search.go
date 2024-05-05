package track

import (
	"context"
	"fmt"

	"github.com/zuzuka28/music_land_api/internal/model"
	"xorm.io/xorm"
)

type searchRepo struct {
	db *xorm.Engine
}

func newSearchRepository(db *xorm.Engine) *searchRepo {
	return &searchRepo{
		db: db,
	}
}

func (r *searchRepo) Search(ctx context.Context, query *model.TrackSearchQuery) ([]*model.Track, error) {
	var tr []*track

	sess := r.db.Context(ctx)

	if query.Name != "" {
		sess = sess.Where("name = ?", query.Name)
	}

	if err := sess.Find(&tr); err != nil {
		return nil, fmt.Errorf("get track from db: %w", err)
	}

	res := make([]*model.Track, 0, len(tr))
	for i := range tr {
		res = append(res, mapTrackToModel(tr[i]))
	}

	return res, nil
}
