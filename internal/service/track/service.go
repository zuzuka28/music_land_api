package track

import (
	"context"

	"github.com/zuzuka28/music_land_api/internal/model"
	"github.com/zuzuka28/music_land_api/pkg/fs"
)

type Service struct {
	tr Tracer

	cs  *createService
	ds  *deleteService
	fs  *fetchService
	ffs *fetchFileService
	ss  *searchService
}

func NewService(r Repository, fss FileStorage, tr Tracer) *Service {
	return &Service{
		tr:  tr,
		cs:  newCreateService(r, fss),
		ds:  newDeleteService(r, fss),
		fs:  newFetchService(r),
		ffs: newFetchFileService(newFetchService(r), fss),
		ss:  newSearchService(r),
	}
}

func (s *Service) Create(ctx context.Context, cmd *model.TrackCreateCommand) error {
	ctx, span := s.tr.Start(ctx, "Create")
	defer span.End()

	return s.cs.Create(ctx, cmd)
}

func (s *Service) Delete(ctx context.Context, cmd *model.TrackDeleteCommand) error {
	ctx, span := s.tr.Start(ctx, "Delete")
	defer span.End()

	return s.ds.Delete(ctx, cmd)
}

func (s *Service) Fetch(ctx context.Context, query *model.TrackFetchQuery) (*model.Track, error) {
	ctx, span := s.tr.Start(ctx, "Fetch")
	defer span.End()

	return s.fs.Fetch(ctx, query)
}

func (s *Service) FetchFile(ctx context.Context, query *model.TrackFetchQuery) (*fs.File, error) {
	ctx, span := s.tr.Start(ctx, "FetchFile")
	defer span.End()

	return s.ffs.FetchFile(ctx, query)
}

func (s *Service) Search(ctx context.Context, query *model.TrackSearchQuery) ([]*model.Track, error) {
	ctx, span := s.tr.Start(ctx, "Search")
	defer span.End()

	return s.ss.Search(ctx, query)
}
