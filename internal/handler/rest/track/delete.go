package track

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zuzuka28/music_land_api/internal/handler/rest/response"
	"github.com/zuzuka28/music_land_api/internal/model"
)

type deleter interface {
	Delete(ctx context.Context, cmd *model.TrackDeleteCommand) error
}

func makeDeleteHandler(s deleter) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		req, err := parseTrackDeleteCommand(gctx)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, response.Error{Message: err.Error()})
			return
		}

		if err := s.Delete(gctx.Request.Context(), req); err != nil {
			gctx.JSON(http.StatusInternalServerError, response.Error{Message: err.Error()})
			return
		}

		gctx.JSON(http.StatusOK, deleteResponse{
			Status: "OK",
		})
	}
}
