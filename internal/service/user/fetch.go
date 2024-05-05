package user

import (
	"context"
	"fmt"

	"github.com/zuzuka28/music_land_api/internal/model"
)

type fetcher interface {
	Fetch(ctx context.Context, id string) (*model.User, error)
}

type fetchService struct {
	r fetcher
}

func newFetchService(r Repository) *fetchService {
	return &fetchService{
		r: r,
	}
}

func (s *fetchService) Fetch(ctx context.Context, query *model.UserFetchQuery) (*model.User, error) {
	res, err := s.r.Fetch(ctx, query.Nickname)
	if err != nil {
		return nil, fmt.Errorf("fetch user: %w", err)
	}

	return res, nil
}
