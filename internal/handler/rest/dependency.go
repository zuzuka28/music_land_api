package rest

import "github.com/gin-gonic/gin"

type (
	Logger interface {
		Info(msg string, args ...any)
		Warn(msg string, args ...any)
		Error(msg string, args ...any)
	}

	TrackHandler interface {
		Delete(gctx *gin.Context)
		Fetch(gctx *gin.Context)
		Create(gctx *gin.Context)
		Search(gctx *gin.Context)
		FetchFile(gctx *gin.Context)
	}

	UserHandler interface {
		Create(gctx *gin.Context)
	}

	AlbumHandler interface {
		Delete(gctx *gin.Context)
		Fetch(gctx *gin.Context)
		Create(gctx *gin.Context)
		Update(gctx *gin.Context)
	}
)
