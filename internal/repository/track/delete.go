package track

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
	tr := &track{ //nolint:exhaustruct
		UID: id,
	}

	_, err := r.db.Context(ctx).Delete(tr)
	if err != nil {
		return fmt.Errorf("delete track from db: %w", err)
	}

	return nil
}
