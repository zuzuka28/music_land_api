package track

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zuzuka28/music_land_api/internal/handler/rest/response"
	"github.com/zuzuka28/music_land_api/internal/model"
)

type creator interface {
	Create(ctx context.Context, cmd *model.TrackCreateCommand) error
}

func makeCreateHandler(s creator) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		req, err := parseTrackSaveCommand(gctx)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, response.Error{Message: err.Error()})
			return
		}

		if err := s.Create(gctx.Request.Context(), req); err != nil {
			gctx.JSON(http.StatusInternalServerError, response.Error{Message: err.Error()})
			return
		}

		gctx.JSON(http.StatusOK, saveResponse{
			Status: "OK",
		})
	}
}
