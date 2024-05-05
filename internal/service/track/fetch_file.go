package track

import (
	"context"
	"fmt"

	"github.com/zuzuka28/music_land_api/internal/model"
	"github.com/zuzuka28/music_land_api/pkg/fs"
)

type fileFetcher interface {
	FetchFile(ctx context.Context, name string) (*fs.File, error)
}

type fetchFileService struct {
	fs fileFetcher
	f  *fetchService
}

func newFetchFileService(f *fetchService, fss fileFetcher) *fetchFileService {
	return &fetchFileService{
		fs: fss,
		f:  f,
	}
}

func (s *fetchFileService) FetchFile(ctx context.Context, query *model.TrackFetchQuery) (*fs.File, error) {
	item, err := s.f.Fetch(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("fetch track: %w", err)
	}

	res, err := s.fs.FetchFile(ctx, item.FileID)
	if err != nil {
		return nil, fmt.Errorf("fetch file: %w", err)
	}

	return res, nil
}
