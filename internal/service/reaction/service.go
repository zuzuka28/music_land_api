package reaction

import (
	"context"

	"github.com/zuzuka28/music_land_api/internal/model"
)

type Service struct {
	tr Tracer

	cs *createService
	ds *deleteService
	ss *searchService
}

func NewService(r Repository, tr Tracer) *Service {
	return &Service{
		tr: tr,
		cs: newCreateService(r),
		ds: newDeleteService(r),
		ss: newSearchService(r),
	}
}

func (s *Service) Create(ctx context.Context, cmd *model.ReactionCreateCommand) error {
	ctx, span := s.tr.Start(ctx, "Create")
	defer span.End()

	return s.cs.Create(ctx, cmd)
}

func (s *Service) Delete(ctx context.Context, cmd *model.ReactionDeleteCommand) error {
	ctx, span := s.tr.Start(ctx, "Delete")
	defer span.End()

	return s.ds.Delete(ctx, cmd)
}

func (s *Service) Search(
	ctx context.Context,
	query *model.ReactionSearchQuery,
) ([]*model.Reaction, error) {
	ctx, span := s.tr.Start(ctx, "Search")
	defer span.End()

	return s.ss.Search(ctx, query)
}
