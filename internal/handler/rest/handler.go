package rest

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zuzuka28/music_land_api/internal/handler/rest/middleware/logging"
)

type Handler struct {
	eng *gin.Engine
}

func NewHandler(
	userHandler UserHandler,
	trackHandler TrackHandler,
	authMiddleware gin.HandlerFunc,
	l Logger,
) *Handler {
	eng := gin.New()

	api := eng.Group("/api")
	api.Use(
		logging.InjectRequestID(),
		logging.Middleware(l),
		cors.Default(),
	)

	userAPI := api.Group("/user")
	userAPI.POST("", userHandler.Create)

	trackAPI := api.Group("/track")
	trackAPI.Use(authMiddleware)
	trackAPI.GET("/search", trackHandler.Search)
	trackAPI.GET("/:trackID", trackHandler.Fetch)
	trackAPI.DELETE("/:trackID", trackHandler.Delete)
	trackAPI.GET("/:trackID/download", trackHandler.FetchFile)
	trackAPI.POST("", trackHandler.Create)

	return &Handler{
		eng: eng,
	}
}

func (h *Handler) Run(addr ...string) error {
	if err := h.eng.Run(addr...); err != nil {
		return fmt.Errorf("run REST API: %w", err)
	}

	return nil
}
