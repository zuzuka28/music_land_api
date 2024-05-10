package album

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zuzuka28/music_land_api/internal/handler/rest/response"
	"github.com/zuzuka28/music_land_api/internal/model"
)

type fetcher interface {
	Fetch(ctx context.Context, query *model.AlbumFetchQuery) (*model.Album, error)
}

func makeFetchHandler(s fetcher) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		req, err := parseAlbumFetchQuery(gctx)
		if err != nil {
			gctx.JSON(response.NewError(err))
			return
		}

		res, err := s.Fetch(gctx.Request.Context(), req)
		if err != nil {
			gctx.JSON(response.NewError(err))
			return
		}

		gctx.JSON(http.StatusOK, mapFetchResponse(res))
	}
}
