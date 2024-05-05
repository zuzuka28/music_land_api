package rest

import "github.com/gin-gonic/gin"

type (
	TrackHandler interface {
		Delete(gctx *gin.Context)
		Fetch(gctx *gin.Context)
		Create(gctx *gin.Context)
		Search(gctx *gin.Context)
	}

	UserHandler interface {
		Create(gctx *gin.Context)
	}
)
