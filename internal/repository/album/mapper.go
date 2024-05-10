package album

import (
	"time"

	"github.com/zuzuka28/music_land_api/internal/model"
)

func mapAlbumToInternal(in *model.Album) *album {
	if in == nil {
		return nil
	}

	return &album{
		UID:     in.ID,
		Name:    in.Name,
		OwnerID: in.OwnerID,
		Created: time.Time{},
		Updated: time.Time{},
	}
}

func mapAlbumItems(in *model.Album) []*albumItem {
	if in == nil {
		return nil
	}

	items := make([]*albumItem, 0, len(in.TrackIDs))
	for _, v := range in.TrackIDs {
		items = append(items, &albumItem{
			AlbumID: in.ID,
			TrackID: v,
			Created: time.Time{},
			Updated: time.Time{},
		})
	}

	return items
}

func mapAlbumToModel(in *album) *model.Album {
	if in == nil {
		return nil
	}

	return &model.Album{
		ID:       in.UID,
		Name:     in.Name,
		OwnerID:  in.OwnerID,
		TrackIDs: nil,
	}
}

func collectTrackIDs(in []*albumItem) []string {
	ids := make([]string, 0, len(in))
	for i := range in {
		ids = append(ids, in[i].TrackID)
	}

	return ids
}
