package user

import (
	"context"
	"fmt"

	"github.com/zuzuka28/music_land_api/internal/model"
	"xorm.io/xorm"
)

type Repository struct {
	*createRepo
	*fetchRepo
}

func NewRepository(db *xorm.Engine) (*Repository, error) {
	if err := db.Sync(new(user)); err != nil {
		return nil, fmt.Errorf("sync user: %w", err)
	}

	return &Repository{
		createRepo: newCreateRepository(db),
		fetchRepo:  newFetchRepository(db),
	}, nil
}

func (r *Repository) Create(ctx context.Context, item *model.User) error {
	return r.createRepo.Create(ctx, item)
}

func (r *Repository) Fetch(ctx context.Context, id string) (*model.User, error) {
	return r.fetchRepo.Fetch(ctx, id)
}
