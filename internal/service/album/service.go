package album

import (
	"context"

	"github.com/zuzuka28/music_land_api/internal/model"
)

type Service struct {
	tr Tracer

	cs *createService
	ds *deleteService
	fs *fetchService
	us *updateService
}

func NewService(r Repository, tr Tracer) *Service {
	return &Service{
		tr: tr,
		cs: newCreateService(r),
		ds: newDeleteService(r),
		fs: newFetchService(r),
		us: newUpdateService(r),
	}
}

func (s *Service) Create(ctx context.Context, cmd *model.AlbumCreateCommand) error {
	ctx, span := s.tr.Start(ctx, "Create")
	defer span.End()

	return s.cs.Create(ctx, cmd)
}

func (s *Service) Delete(ctx context.Context, cmd *model.AlbumDeleteCommand) error {
	ctx, span := s.tr.Start(ctx, "Delete")
	defer span.End()

	return s.ds.Delete(ctx, cmd)
}

func (s *Service) Fetch(ctx context.Context, query *model.AlbumFetchQuery) (*model.Album, error) {
	ctx, span := s.tr.Start(ctx, "Fetch")
	defer span.End()

	return s.fs.Fetch(ctx, query)
}

func (s *Service) Update(ctx context.Context, cmd *model.AlbumUpdateCommand) error {
	ctx, span := s.tr.Start(ctx, "Update")
	defer span.End()

	return s.us.Update(ctx, cmd)
}
