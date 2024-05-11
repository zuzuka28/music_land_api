package album

import (
	"context"
	"fmt"

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

func (r *deleteRepo) Delete(ctx context.Context, id string) error {
	_, err := r.db.Transaction(func(sess *xorm.Session) (any, error) {
		sess = sess.Context(ctx)

		if _, err := sess.Delete(&album{UID: id}); err != nil { //nolint:exhaustruct
			return nil, fmt.Errorf("delete album in db: %w", err)
		}

		if _, err := sess.Delete(&albumItem{AlbumID: id}); err != nil { //nolint:exhaustruct
			return nil, fmt.Errorf("delete album items in db: %w", err)
		}

		return nil, nil //nolint:nilnil
	})
	if err != nil {
		return fmt.Errorf("delete album: %w", err)
	}

	return nil
}
