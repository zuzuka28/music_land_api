package track

import (
	"context"
	"fmt"

	"github.com/zuzuka28/music_land_api/internal/model"
)

type fileDeleter interface {
	DeleteFile(ctx context.Context, name string) error
}

type fetcherDeleter interface {
	Fetch(ctx context.Context, id string) (*model.Track, error)
	Delete(ctx context.Context, id string) error
}

type deleteService struct {
	fs fileDeleter
	r  fetcherDeleter
}

func newDeleteService(r Repository, fs fileDeleter) *deleteService {
	return &deleteService{
		fs: fs,
		r:  r,
	}
}

func (s *deleteService) Delete(ctx context.Context, cmd *model.TrackDeleteCommand) error {
	item, err := s.r.Fetch(ctx, cmd.ID)
	if err != nil {
		return fmt.Errorf("fetch track: %w", err)
	}

	if err := s.fs.DeleteFile(ctx, item.FileID); err != nil {
		return fmt.Errorf("delete track from file storage: %w", err)
	}

	if err := s.r.Delete(ctx, cmd.ID); err != nil {
		return fmt.Errorf("delete track: %w", err)
	}

	return nil
}
