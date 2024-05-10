package track

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/zuzuka28/music_land_api/internal/model"
)

func parseTrackDeleteCommand(gctx *gin.Context) (*model.TrackDeleteCommand, error) {
	trackID := gctx.Param("trackID")

	if trackID == "" {
		return nil, fmt.Errorf("%w: empty trackID", model.ErrNotValid)
	}

	return &model.TrackDeleteCommand{
		ID: trackID,
	}, nil
}

// FIXME: do not read file, pass it within io.Reader.
func parseTrackSaveCommand(gctx *gin.Context) (*model.TrackCreateCommand, error) {
	req := new(saveRequest)
	if err := gctx.ShouldBind(req); err != nil {
		return nil, fmt.Errorf("bind save request: %w", err)
	}

	f, _ := req.Attachment.Open()
	data, _ := io.ReadAll(f)

	return &model.TrackCreateCommand{
		Name:      req.Name,
		Author:    req.Author,
		TrackData: data,
	}, nil
}

func parseTrackFetchQuery(gctx *gin.Context) (*model.TrackFetchQuery, error) {
	trackID := gctx.Param("trackID")

	if trackID == "" {
		return nil, fmt.Errorf("%w: empty trackID", model.ErrNotValid)
	}

	return &model.TrackFetchQuery{
		ID: trackID,
	}, nil
}

func mapFetchResponse(in *model.Track) *fetchResponse {
	if in == nil {
		return nil
	}

	return &fetchResponse{
		ID:     in.ID,
		Name:   in.Name,
		Author: in.Author,
		FileID: in.FileID,
	}
}

func parseTrackSearchQuery(gctx *gin.Context) (*model.TrackSearchQuery, error) {
	req := new(searchRequest)
	if err := gctx.ShouldBind(req); err != nil {
		return nil, fmt.Errorf("bind search request: %w", err)
	}

	return &model.TrackSearchQuery{
		Name: req.Name,
	}, nil
}

func mapSearchResponse(in []*model.Track) *searchResponse {
	items := make([]*track, 0, len(in))

	for i := range in {
		items = append(items, mapTrackToResponse(in[i]))
	}

	return &searchResponse{
		Items: items,
	}
}

func mapTrackToResponse(in *model.Track) *track {
	if in == nil {
		return nil
	}

	return &track{
		ID:     in.ID,
		Name:   in.Name,
		Author: in.Author,
	}
}
