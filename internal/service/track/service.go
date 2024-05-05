package track

import (
	"context"

	"github.com/zuzuka28/music_land_api/internal/model"
	"github.com/zuzuka28/music_land_api/pkg/fs"
)

type Service struct {
	*createService
	*deleteService
	*fetchService
	*fetchFileService
	*searchService
}

func NewService(r Repository, fss FileStorage) *Service {
	return &Service{
		createService:    newCreateService(r, fss),
		deleteService:    newDeleteService(r, fss),
		fetchService:     newFetchService(r),
		fetchFileService: newFetchFileService(newFetchService(r), fss),
		searchService:    newSearchService(r),
	}
}

func (s *Service) Create(ctx context.Context, cmd *model.TrackCreateCommand) error {
	return s.createService.Create(ctx, cmd)
}

func (s *Service) Delete(ctx context.Context, cmd *model.TrackDeleteCommand) error {
	return s.deleteService.Delete(ctx, cmd)
}

func (s *Service) Fetch(ctx context.Context, query *model.TrackFetchQuery) (*model.Track, error) {
	return s.fetchService.Fetch(ctx, query)
}

func (s *Service) FetchFile(ctx context.Context, query *model.TrackFetchQuery) (*fs.File, error) {
	return s.fetchFileService.FetchFile(ctx, query)
}

func (s *Service) Search(ctx context.Context, query *model.TrackSearchQuery) ([]*model.Track, error) {
	return s.searchService.Search(ctx, query)
}
