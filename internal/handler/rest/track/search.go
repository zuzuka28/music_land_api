package track

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zuzuka28/music_land_api/internal/handler/rest/response"
	"github.com/zuzuka28/music_land_api/internal/model"
)

type searcher interface {
	Search(ctx context.Context, query *model.TrackSearchQuery) ([]*model.Track, error)
}

func makeSearchHandler(s searcher) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		req, err := parseTrachSearchQuery(gctx)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, response.Error{Message: err.Error()})
			return
		}

		res, err := s.Search(gctx.Request.Context(), req)
		if err != nil {
			gctx.JSON(http.StatusInternalServerError, response.Error{Message: err.Error()})
			return
		}

		gctx.JSON(http.StatusOK, mapSearchResponse(res))
	}
}
