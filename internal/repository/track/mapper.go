package track

import (
	"time"

	"github.com/google/uuid"
	"github.com/zuzuka28/music_land_api/internal/model"
)

func mapTrackToModel(in *track) *model.Track {
	if in == nil {
		return nil
	}

	return &model.Track{
		ID:     in.UID,
		Name:   in.Name,
		Author: in.Author,
		FileID: in.FileID,
	}
}

func mapTrackToInternal(in *model.Track) *track {
	if in == nil {
		return nil
	}

	id := uuid.New().String()

	return &track{
		UID:     id,
		Name:    in.Name,
		Author:  in.Author,
		FileID:  in.FileID,
		Created: time.Time{},
		Updated: time.Time{},
	}
}
