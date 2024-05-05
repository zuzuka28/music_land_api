package user

import (
	"context"

	"github.com/zuzuka28/music_land_api/internal/model"
)

type Service struct {
	*createService
	*fetchService
}

func NewService(r Repository) *Service {
	return &Service{
		createService: newCreateService(r),
		fetchService:  newFetchService(r),
	}
}

func (s *Service) Create(ctx context.Context, cmd *model.UserCreateCommand) error {
	return s.createService.Create(ctx, cmd)
}

func (s *Service) Fetch(ctx context.Context, query *model.UserFetchQuery) (*model.User, error) {
	return s.fetchService.Fetch(ctx, query)
}
