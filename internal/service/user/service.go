package user

import (
	"context"

	"github.com/zuzuka28/music_land_api/internal/model"
)

type Service struct {
	tr Tracer

	cs *createService
	fs *fetchService
}

func NewService(r Repository, tr Tracer) *Service {
	return &Service{
		tr: tr,
		cs: newCreateService(r),
		fs: newFetchService(r),
	}
}

func (s *Service) Create(ctx context.Context, cmd *model.UserCreateCommand) error {
	ctx, span := s.tr.Start(ctx, "Create")
	defer span.End()

	return s.cs.Create(ctx, cmd)
}

func (s *Service) Fetch(ctx context.Context, query *model.UserFetchQuery) (*model.User, error) {
	ctx, span := s.tr.Start(ctx, "Fetch")
	defer span.End()

	return s.fs.Fetch(ctx, query)
}
