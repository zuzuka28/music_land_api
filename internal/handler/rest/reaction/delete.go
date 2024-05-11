package reaction

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zuzuka28/music_land_api/internal/handler/rest/response"
	"github.com/zuzuka28/music_land_api/internal/model"
)

type deleter interface {
	Delete(ctx context.Context, cmd *model.ReactionDeleteCommand) error
}

func makeDeleteHandler(s deleter) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		req, err := parseReactionDeleteCommand(gctx)
		if err != nil {
			gctx.JSON(response.NewError(err))
			return
		}

		if err := s.Delete(gctx.Request.Context(), req); err != nil {
			gctx.JSON(response.NewError(err))
			return
		}

		gctx.JSON(http.StatusOK, deleteResponse{
			Status: "OK",
		})
	}
}
