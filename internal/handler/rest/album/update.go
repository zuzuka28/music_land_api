package album

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zuzuka28/music_land_api/internal/handler/rest/response"
	"github.com/zuzuka28/music_land_api/internal/model"
)

type updater interface {
	Update(ctx context.Context, cmd *model.AlbumUpdateCommand) error
}

func makeUpdateHandler(s updater) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		req, err := parseAlbumUpdateCommand(gctx)
		if err != nil {
			gctx.JSON(response.NewError(err))
			return
		}

		if err := s.Update(gctx.Request.Context(), req); err != nil {
			gctx.JSON(response.NewError(err))
			return
		}

		gctx.JSON(http.StatusOK, updateResponse{
			Status: "OK",
		})
	}
}
