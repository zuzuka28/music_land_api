package logging

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	httpLast2xxStatus  = 299
	httpFirst5xxStatus = 500
)

func Middleware(l Logger) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		l.Info("incomming request",
			"requestID", gctx.GetString(RequestIDKey()),
			"url", gctx.Request.URL.String(),
			gctx.Writer.Status,
		)

		start := time.Now()

		gctx.Next()

		log := logFuncByStatus(l, gctx.Writer.Size())

		log("request processed",
			"requestID", gctx.GetString(RequestIDKey()),
			"url", gctx.Request.URL.String(),
			"elapsedTime", time.Since(start).String(),
			"size", gctx.Writer.Size(),
			"status", gctx.Writer.Status(),
		)
	}
}

func InjectRequestID() gin.HandlerFunc {
	return func(gctx *gin.Context) {
		gctx.Set(RequestIDKey(), uuid.New().String())
		gctx.Next()
	}
}

func logFuncByStatus(l Logger, code int) func(msg string, args ...any) {
	switch {
	case code <= httpLast2xxStatus:
		return l.Info

	case code > httpLast2xxStatus && code < httpFirst5xxStatus:
		return l.Warn

	case code >= httpFirst5xxStatus:
		return l.Error

	default:
		return l.Info
	}
}
