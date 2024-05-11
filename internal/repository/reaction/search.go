package reaction

import (
	"context"
	"fmt"

	"github.com/zuzuka28/music_land_api/internal/model"
	"xorm.io/xorm"
)

type searchRepo struct {
	db *xorm.Engine
}

func newSearchRepository(db *xorm.Engine) *searchRepo {
	return &searchRepo{
		db: db,
	}
}

func (r *searchRepo) Search(
	ctx context.Context,
	query *model.ReactionSearchQuery,
) ([]*model.Reaction, error) {
	var react []*reaction

	if err := r.db.Context(ctx).Find(&react, &reaction{ //nolint:exhaustruct
		UserID:             query.UserID,
		ReactionType:       string(query.ReactionType),
		ReactionTargetID:   query.ReactionTargetID,
		ReactionTargetType: string(query.ReactionType),
	}); err != nil {
		return nil, fmt.Errorf("get reaction from db: %w", err)
	}

	return mapReactionsToModel(react), nil
}
