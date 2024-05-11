package album

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

func (r *fetchRepo) Fetch(ctx context.Context, id string) (*model.Album, error) {
	var (
		alb = &album{ //nolint:exhaustruct
			UID: id,
		}
		tracks []*albumItem
	)

	_, err := r.db.Transaction(func(sess *xorm.Session) (any, error) {
		sess = sess.Context(ctx)

		has, err := sess.Get(alb)
		if err != nil {
			return nil, fmt.Errorf("delete album in db: %w", err)
		}

		if !has {
			return nil, fmt.Errorf("%w: get album %s in db", model.ErrNotFound, id)
		}

		if err := sess.Find(tracks, &albumItem{AlbumID: id}); err != nil { //nolint:exhaustruct
			return nil, fmt.Errorf("find album items in db: %w", err)
		}

		return nil, nil //nolint:nilnil
	})
	if err != nil {
		return nil, fmt.Errorf("fetch album: %w", err)
	}

	res := mapAlbumToModel(alb)
	res.TrackIDs = collectTrackIDs(tracks)

	return res, nil
}
