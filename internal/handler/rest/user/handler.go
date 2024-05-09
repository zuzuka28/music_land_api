package user

import "github.com/gin-gonic/gin"

type Handler struct {
	tr Tracer

	create gin.HandlerFunc
}

func NewHandler(s Service, tr Tracer) *Handler {
	return &Handler{
		tr:     tr,
		create: makeCreateHandler(s),
	}
}

func (h *Handler) Create(gctx *gin.Context) {
	ctx, span := h.tr.Start(gctx.Request.Context(), "Create")
	defer span.End()

	gctx.Request = gctx.Request.WithContext(ctx)

	h.create(gctx)
}
