package reaction

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
	sr *searchRepo
}

func NewRepository(db *xorm.Engine, tr Tracer) (*Repository, error) {
	if err := db.Sync(new(reaction)); err != nil {
		return nil, fmt.Errorf("sync reaction: %w", err)
	}

	return &Repository{
		tr: tr,
		cr: newCreateRepository(db),
		dr: newDeleteRepository(db),
		sr: newSearchRepository(db),
	}, nil
}

func (r *Repository) Create(ctx context.Context, item *model.Reaction) error {
	ctx, span := r.tr.Start(ctx, "Create")
	defer span.End()

	return r.cr.Create(ctx, item)
}

func (r *Repository) Delete(ctx context.Context, cmd *model.ReactionDeleteCommand) error {
	ctx, span := r.tr.Start(ctx, "Delete")
	defer span.End()

	return r.dr.Delete(ctx, cmd)
}

func (r *Repository) Search(
	ctx context.Context,
	query *model.ReactionSearchQuery,
) ([]*model.Reaction, error) {
	ctx, span := r.tr.Start(ctx, "Search")
	defer span.End()

	return r.sr.Search(ctx, query)
}
