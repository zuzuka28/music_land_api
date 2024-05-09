package user

import (
	"context"
	"fmt"

	"github.com/zuzuka28/music_land_api/internal/model"
	"xorm.io/xorm"
)

type Repository struct {
	tr Tracer

	cr *createRepo
	fr *fetchRepo
}

func NewRepository(db *xorm.Engine, tr Tracer) (*Repository, error) {
	if err := db.Sync(new(user)); err != nil {
		return nil, fmt.Errorf("sync user: %w", err)
	}

	return &Repository{
		tr: tr,
		cr: newCreateRepository(db),
		fr: newFetchRepository(db),
	}, nil
}

func (r *Repository) Create(ctx context.Context, item *model.User) error {
	ctx, span := r.tr.Start(ctx, "Create")
	defer span.End()

	return r.cr.Create(ctx, item)
}

func (r *Repository) Fetch(ctx context.Context, id string) (*model.User, error) {
	ctx, span := r.tr.Start(ctx, "Fetch")
	defer span.End()

	return r.fr.Fetch(ctx, id)
}
