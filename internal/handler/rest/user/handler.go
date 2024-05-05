package user

import "github.com/gin-gonic/gin"

type Handler struct {
	create gin.HandlerFunc
}

func NewHandler(s Service) *Handler {
	return &Handler{
		create: makeCreateHandler(s),
	}
}

func (h *Handler) Create(gctx *gin.Context) {
	h.create(gctx)
}
