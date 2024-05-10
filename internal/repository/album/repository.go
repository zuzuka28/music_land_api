package album

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
	ur *updateRepo
}

func NewRepository(db *xorm.Engine, tr Tracer) (*Repository, error) {
	if err := db.Sync(new(album), new(albumItem)); err != nil {
		return nil, fmt.Errorf("sync album: %w", err)
	}

	return &Repository{
		tr: tr,
		cr: newCreateRepository(db),
		dr: newDeleteRepository(db),
		fr: newFetchRepository(db),
		ur: newUpdateRepository(db),
	}, nil
}

func (r *Repository) Create(ctx context.Context, item *model.Album) error {
	ctx, span := r.tr.Start(ctx, "Create")
	defer span.End()

	return r.cr.Create(ctx, item)
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	ctx, span := r.tr.Start(ctx, "Delete")
	defer span.End()

	return r.dr.Delete(ctx, id)
}

func (r *Repository) Fetch(ctx context.Context, id string) (*model.Album, error) {
	ctx, span := r.tr.Start(ctx, "Fetch")
	defer span.End()

	return r.fr.Fetch(ctx, id)
}

func (r *Repository) Update(ctx context.Context, cmd *model.AlbumUpdateCommand) error {
	ctx, span := r.tr.Start(ctx, "Update")
	defer span.End()

	return r.ur.Update(ctx, cmd)
}
