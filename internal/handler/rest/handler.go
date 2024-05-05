package rest

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	eng *gin.Engine
}

func NewHandler(
	userHandler UserHandler,
	trackHandler TrackHandler,
	authMiddleware gin.HandlerFunc,
) *Handler {
	eng := gin.New()

	api := eng.Group("/api")
	api.Use(cors.Default())

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
