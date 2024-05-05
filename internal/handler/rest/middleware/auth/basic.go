package auth

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zuzuka28/music_land_api/internal/handler/rest/response"
	"github.com/zuzuka28/music_land_api/internal/model"
)

func BasicMiddleware(s Service) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		l, p, ok := gctx.Request.BasicAuth()

		if !ok {
			gctx.JSON(http.StatusForbidden, response.Error{
				Message: http.StatusText(http.StatusForbidden),
			})
			gctx.Abort()

			return
		}

		usr, err := s.Authenticate(gctx.Request.Context(), &model.AuthCommand{
			Login:    l,
			Password: p,
		})
		if err != nil {
			slog.Warn(err.Error())
			gctx.JSON(http.StatusUnauthorized, response.Error{
				Message: http.StatusText(http.StatusUnauthorized),
			})
			gctx.Abort()

			return
		}

		gctx.Set(IdentityKey(), usr)

		gctx.Next()
	}
}
