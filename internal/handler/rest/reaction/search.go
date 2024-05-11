package reaction

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zuzuka28/music_land_api/internal/handler/rest/response"
	"github.com/zuzuka28/music_land_api/internal/model"
)

type searcher interface {
	Search(ctx context.Context, query *model.ReactionSearchQuery) ([]*model.Reaction, error)
}

func makeSearchHandler(s searcher) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		req, err := parseReactionSearchQuery(gctx)
		if err != nil {
			gctx.JSON(response.NewError(err))
			return
		}

		res, err := s.Search(gctx.Request.Context(), req)
		if err != nil {
			gctx.JSON(response.NewError(err))
			return
		}

		gctx.JSON(http.StatusOK, mapSearchResponse(res))
	}
}
