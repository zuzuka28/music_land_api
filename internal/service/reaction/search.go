package reaction

import (
	"context"
	"fmt"

	"github.com/zuzuka28/music_land_api/internal/model"
)

type searcher interface {
	Search(ctx context.Context, query *model.ReactionSearchQuery) ([]*model.Reaction, error)
}

type searchService struct {
	r searcher
}

func newSearchService(r searcher) *searchService {
	return &searchService{
		r: r,
	}
}

func (s *searchService) Search(
	ctx context.Context,
	query *model.ReactionSearchQuery,
) ([]*model.Reaction, error) {
	res, err := s.r.Search(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("search reactions: %w", err)
	}

	return res, nil
}
