package track

import (
	"context"
	"fmt"

	"github.com/zuzuka28/music_land_api/internal/model"
	"xorm.io/xorm"
)

type Repository struct {
	tr Tracer

	cr *createRepo
	dr *deleteRepo
	fr *fetchRepo
	sr *searchRepo
}

func NewRepository(db *xorm.Engine, tr Tracer) (*Repository, error) {
	if err := db.Sync(new(track)); err != nil {
		return nil, fmt.Errorf("sync track: %w", err)
	}

	return &Repository{
		tr: tr,
		cr: newCreateRepository(db),
		dr: newDeleteRepository(db),
		fr: newFetchRepository(db),
		sr: newSearchRepository(db),
	}, nil
}

func (r *Repository) Create(ctx context.Context, item *model.Track) error {
	ctx, span := r.tr.Start(ctx, "Create")
	defer span.End()

	return r.cr.Create(ctx, item)
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	ctx, span := r.tr.Start(ctx, "Delete")
	defer span.End()

	return r.dr.Delete(ctx, id)
}

func (r *Repository) Fetch(ctx context.Context, id string) (*model.Track, error) {
	ctx, span := r.tr.Start(ctx, "Fetch")
	defer span.End()

	return r.fr.Fetch(ctx, id)
}

func (r *Repository) Search(ctx context.Context, query *model.TrackSearchQuery) ([]*model.Track, error) {
	ctx, span := r.tr.Start(ctx, "Search")
	defer span.End()

	return r.sr.Search(ctx, query)
}
