package track

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zuzuka28/music_land_api/internal/handler/rest/response"
	"github.com/zuzuka28/music_land_api/internal/model"
	"github.com/zuzuka28/music_land_api/pkg/fs"
)

type fileFetcher interface {
	FetchFile(ctx context.Context, query *model.TrackFetchQuery) (*fs.File, error)
}

func makeFetchFileHandler(s fileFetcher) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		req, err := parseTrackFetchQuery(gctx)
		if err != nil {
			gctx.JSON(http.StatusBadRequest, response.Error{Message: err.Error()})
			return
		}

		res, err := s.FetchFile(gctx.Request.Context(), req)
		if err != nil {
			gctx.JSON(http.StatusInternalServerError, response.Error{Message: err.Error()})
			return
		}

		gctx.DataFromReader(
			http.StatusOK,
			res.Size,
			"application/gzip",
			res.Reader,
			map[string]string{
				"Content-Disposition": `attachment; filename=` + res.Name,
			},
		)
	}
}
