package reaction

import "github.com/gin-gonic/gin"

type Handler struct {
	tr Tracer

	delete gin.HandlerFunc
	create gin.HandlerFunc
	search gin.HandlerFunc
}

func NewHandler(s Service, tr Tracer) *Handler {
	return &Handler{
		tr:     tr,
		delete: makeDeleteHandler(s),
		create: makeCreateHandler(s),
		search: makeSearchHandler(s),
	}
}

func (h *Handler) Delete(gctx *gin.Context) {
	ctx, span := h.tr.Start(gctx.Request.Context(), "Delete")
	defer span.End()

	gctx.Request = gctx.Request.WithContext(ctx)

	h.delete(gctx)
}

func (h *Handler) Create(gctx *gin.Context) {
	ctx, span := h.tr.Start(gctx.Request.Context(), "Create")
	defer span.End()

	gctx.Request = gctx.Request.WithContext(ctx)

	h.create(gctx)
}

func (h *Handler) Search(gctx *gin.Context) {
	ctx, span := h.tr.Start(gctx.Request.Context(), "Search")
	defer span.End()

	gctx.Request = gctx.Request.WithContext(ctx)

	h.search(gctx)
}
