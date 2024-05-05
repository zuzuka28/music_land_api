package track

import (
	"context"
	"fmt"

	"github.com/zuzuka28/music_land_api/internal/model"
)

type fetcher interface {
	Fetch(ctx context.Context, id string) (*model.Track, error)
}

type fetchService struct {
	r fetcher
}

func newFetchService(r Repository) *fetchService {
	return &fetchService{
		r: r,
	}
}

func (s *fetchService) Fetch(ctx context.Context, query *model.TrackFetchQuery) (*model.Track, error) {
	res, err := s.r.Fetch(ctx, query.ID)
	if err != nil {
		return nil, fmt.Errorf("fetch track: %w", err)
	}

	return res, nil
}
