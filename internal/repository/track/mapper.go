package track

import (
	"strconv"
	"time"

	"github.com/zuzuka28/music_land_api/internal/model"
)

func mapTrackToModel(in *track) *model.Track {
	if in == nil {
		return nil
	}

	return &model.Track{
		ID:     strconv.Itoa(int(in.ID)),
		Name:   in.Name,
		Author: in.Author,
		FileID: in.FileID,
	}
}

func mapTrackToInternal(in *model.Track) *track {
	if in == nil {
		return nil
	}

	return &track{
		ID:      0,
		Name:    in.Name,
		Author:  in.Author,
		FileID:  in.FileID,
		Created: time.Time{},
		Updated: time.Time{},
	}
}
