package album

import (
	"context"
	"fmt"

	"github.com/zuzuka28/music_land_api/internal/model"
	"xorm.io/xorm"
)

type updateRepo struct {
	db *xorm.Engine
}

func newUpdateRepository(db *xorm.Engine) *updateRepo {
	return &updateRepo{
		db: db,
	}
}

func (r *updateRepo) Update(ctx context.Context, cmd *model.AlbumUpdateCommand) error {
	var (
		alb = &album{ //nolint:exhaustruct
			UID: cmd.AlbumID,
		}
		upditems = mapAlbumItems(&model.Album{ //nolint:exhaustruct
			ID:       cmd.AlbumID,
			TrackIDs: cmd.TrackIDs,
		})
	)

	_, err := r.db.Transaction(func(sess *xorm.Session) (any, error) {
		sess = sess.Context(ctx)

		if _, err := sess.Update(alb, alb); err != nil {
			return nil, fmt.Errorf("upd album in db: %w", err)
		}

		if _, err := sess.Delete(&albumItem{AlbumID: cmd.AlbumID}); err != nil { //nolint: exhaustruct
			return nil, fmt.Errorf("delete album items in db: %w", err)
		}

		if _, err := sess.Insert(upditems); err != nil {
			return nil, fmt.Errorf("insert album items in db: %w", err)
		}

		return nil, nil //nolint:nilnil
	})
	if err != nil {
		return fmt.Errorf("update album: %w", err)
	}

	return nil
}
