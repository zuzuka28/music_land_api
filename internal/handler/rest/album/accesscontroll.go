package album

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zuzuka28/music_land_api/internal/handler/rest/response"
	"github.com/zuzuka28/music_land_api/internal/model"
)

// albumActionAccessControl - checks OwnerID for album because
// actions with albums should be binded to owners.
func albumActionAccessControl(s fetcher) func(gin.HandlerFunc) gin.HandlerFunc {
	return func(handl gin.HandlerFunc) gin.HandlerFunc {
		return func(gctx *gin.Context) {
			usr, err := parseUserInfo(gctx)
			if err != nil {
				gctx.JSON(response.NewError(err))
				gctx.Abort()

				return
			}

			alb, err := fetchTargetAlbum(s, gctx)
			if err != nil {
				gctx.JSON(response.NewError(err))
				gctx.Abort()

				return
			}

			if alb.OwnerID != usr.ID {
				gctx.JSON(response.NewError(fmt.Errorf("%w: invalid albumID", model.ErrNotValid)))
				gctx.Abort()
			}

			handl(gctx)
		}
	}
}

func fetchTargetAlbum(s fetcher, gctx *gin.Context) (*model.Album, error) {
	albumID := gctx.Param("albumID")

	if albumID == "" {
		return nil, fmt.Errorf("%w: empty albumID", model.ErrNotValid)
	}

	alb, err := s.Fetch(gctx.Request.Context(), &model.AlbumFetchQuery{
		AlbumID: albumID,
	})
	if err != nil {
		return nil, fmt.Errorf("fetch album: %w", err)
	}

	return alb, nil
}
