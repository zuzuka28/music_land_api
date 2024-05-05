package track

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/zuzuka28/music_land_api/internal/model"
)

type fileSaver interface {
	SaveFile(ctx context.Context, name string, data []byte) error
}

type creator interface {
	Create(ctx context.Context, item *model.Track) error
}

type createService struct {
	fs fileSaver
	r  creator
}

func newCreateService(r creator, fs fileSaver) *createService {
	return &createService{
		fs: fs,
		r:  r,
	}
}

func (s *createService) Create(ctx context.Context, cmd *model.TrackCreateCommand) error {
	fname := sha256Content(cmd.TrackData)

	if err := s.fs.SaveFile(ctx, fname, cmd.TrackData); err != nil {
		return fmt.Errorf("save to file storage: %w", err)
	}

	if err := s.r.Create(ctx, &model.Track{
		ID:     "",
		Name:   cmd.Name,
		Author: cmd.Author,
		FileID: fname,
	}); err != nil {
		return fmt.Errorf("create track: %w", err)
	}

	return nil
}

func sha256Content(in []byte) string {
	hash := sha256.New()
	_, _ = hash.Write(in)

	return fmt.Sprintf("%+x", hash.Sum(nil))
}
