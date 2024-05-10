package album

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zuzuka28/music_land_api/internal/handler/rest/middleware/auth"
	"github.com/zuzuka28/music_land_api/internal/model"
)

func parseAlbumDeleteCommand(gctx *gin.Context) (*model.AlbumDeleteCommand, error) {
	albumID := gctx.Param("albumID")

	if albumID == "" {
		return nil, fmt.Errorf("%w: empty albumID", model.ErrNotValid)
	}

	return &model.AlbumDeleteCommand{
		AlbumID: albumID,
	}, nil
}

func parseAlbumCreateCommand(gctx *gin.Context) (*model.AlbumCreateCommand, error) {
	req := new(createRequest)
	if err := gctx.ShouldBind(req); err != nil {
		return nil, fmt.Errorf("bind create request: %w", err)
	}

	usr, err := parseUserInfo(gctx)
	if err != nil {
		return nil, fmt.Errorf("parse user info: %w", err)
	}

	return &model.AlbumCreateCommand{
		OwnerID:  usr.ID,
		Name:     req.Name,
		TrackIDs: req.TrackIDs,
	}, nil
}

func parseAlbumUpdateCommand(gctx *gin.Context) (*model.AlbumUpdateCommand, error) {
	albumID := gctx.Param("albumID")

	if albumID == "" {
		return nil, fmt.Errorf("%w: empty albumID", model.ErrNotValid)
	}

	req := new(updateRequest)
	if err := gctx.ShouldBind(req); err != nil {
		return nil, fmt.Errorf("bind update request: %w", err)
	}

	return &model.AlbumUpdateCommand{
		AlbumID:  albumID,
		TrackIDs: req.TrackIDs,
	}, nil
}

func parseAlbumFetchQuery(gctx *gin.Context) (*model.AlbumFetchQuery, error) {
	albumID := gctx.Param("albumID")

	if albumID == "" {
		return nil, fmt.Errorf("%w: empty albumID", model.ErrNotValid)
	}

	return &model.AlbumFetchQuery{
		AlbumID: albumID,
	}, nil
}

func mapFetchResponse(in *model.Album) *fetchResponse {
	if in == nil {
		return nil
	}

	return &fetchResponse{
		ID:       in.ID,
		Name:     in.Name,
		OwnerID:  in.OwnerID,
		TrackIDs: in.TrackIDs,
	}
}

func parseUserInfo(gctx *gin.Context) (*model.Identity, error) {
	rawusr, has := gctx.Get(auth.IdentityKey())
	if !has {
		return nil, fmt.Errorf("%w: no user info", model.ErrNotValid)
	}

	usr, ok := rawusr.(*model.Identity)
	if !ok {
		return nil, fmt.Errorf("%w: invalid user info", model.ErrNotValid)
	}

	return usr, nil
}
