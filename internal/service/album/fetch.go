package album

import (
	"context"
	"fmt"

	"github.com/zuzuka28/music_land_api/internal/model"
)

type fetcher interface {
	Fetch(ctx context.Context, id string) (*model.Album, error)
}

type fetchService struct {
	r fetcher
}

func newFetchService(r fetcher) *fetchService {
	return &fetchService{
		r: r,
	}
}

func (s *fetchService) Fetch(ctx context.Context, query *model.AlbumFetchQuery) (*model.Album, error) {
	res, err := s.r.Fetch(ctx, query.AlbumID)
	if err != nil {
		return nil, fmt.Errorf("fetch album: %w", err)
	}

	return res, nil
}
