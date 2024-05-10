package album

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zuzuka28/music_land_api/internal/handler/rest/response"
	"github.com/zuzuka28/music_land_api/internal/model"
)

type creator interface {
	Create(ctx context.Context, cmd *model.AlbumCreateCommand) error
}

func makeCreateHandler(s creator) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		req, err := parseAlbumCreateCommand(gctx)
		if err != nil {
			gctx.JSON(response.NewError(err))
			return
		}

		if err := s.Create(gctx.Request.Context(), req); err != nil {
			gctx.JSON(response.NewError(err))
			return
		}

		gctx.JSON(http.StatusOK, createResponse{
			Status: "OK",
		})
	}
}
