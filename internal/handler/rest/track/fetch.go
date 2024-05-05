package track

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zuzuka28/music_land_api/internal/handler/rest/response"
	"github.com/zuzuka28/music_land_api/internal/model"
)

type fetcher interface {
	Fetch(ctx context.Context, query *model.TrackFetchQuery) (*model.Track, error)
}

func makeFetchHandler(s fetcher) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		req, err := parseTrackFetchQuery(gctx)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, response.Error{Message: err.Error()})
			return
		}

		res, err := s.Fetch(gctx.Request.Context(), req)
		if err != nil {
			gctx.JSON(http.StatusInternalServerError, response.Error{Message: err.Error()})
			return
		}

		gctx.JSON(http.StatusOK, mapFetchResponse(res))
	}
}
