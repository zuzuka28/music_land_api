package track

import (
	"context"
	"fmt"

	"github.com/zuzuka28/music_land_api/internal/model"
)

type searcher interface {
	Search(ctx context.Context, query *model.TrackSearchQuery) ([]*model.Track, error)
}

type searchService struct {
	r searcher
}

func newSearchService(r Repository) *searchService {
	return &searchService{
		r: r,
	}
}

func (s *searchService) Search(ctx context.Context, query *model.TrackSearchQuery) ([]*model.Track, error) {
	res, err := s.r.Search(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("search tracks: %w", err)
	}

	return res, nil
}
