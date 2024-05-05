package track

import (
	"context"
	"fmt"

	"github.com/zuzuka28/music_land_api/internal/model"
	"xorm.io/xorm"
)

type Repository struct {
	*createRepo
	*deleteRepo
	*fetchRepo
	*searchRepo
}

func NewRepository(db *xorm.Engine) (*Repository, error) {
	if err := db.Sync(new(track)); err != nil {
		return nil, fmt.Errorf("sync track: %w", err)
	}

	return &Repository{
		createRepo: newCreateRepository(db),
		deleteRepo: newDeleteRepository(db),
		fetchRepo:  newFetchRepository(db),
		searchRepo: newSearchRepository(db),
	}, nil
}

func (r *Repository) Create(ctx context.Context, item *model.Track) error {
	return r.createRepo.Create(ctx, item)
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	return r.deleteRepo.Delete(ctx, id)
}

func (r *Repository) Fetch(ctx context.Context, id string) (*model.Track, error) {
	return r.fetchRepo.Fetch(ctx, id)
}

func (r *Repository) Search(ctx context.Context, query *model.TrackSearchQuery) ([]*model.Track, error) {
	return r.searchRepo.Search(ctx, query)
}
