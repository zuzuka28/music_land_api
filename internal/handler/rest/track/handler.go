package track

import "github.com/gin-gonic/gin"

type Handler struct {
	tr Tracer

	delete    gin.HandlerFunc
	fetch     gin.HandlerFunc
	fetchFile gin.HandlerFunc
	create    gin.HandlerFunc
	search    gin.HandlerFunc
}

func NewHandler(s Service, tr Tracer) *Handler {
	return &Handler{
		tr:        tr,
		delete:    makeDeleteHandler(s),
		fetch:     makeFetchHandler(s),
		fetchFile: makeFetchFileHandler(s),
		create:    makeCreateHandler(s),
		search:    makeSearchHandler(s),
	}
}

func (h *Handler) Delete(gctx *gin.Context) {
	ctx, span := h.tr.Start(gctx.Request.Context(), "Delete")
	defer span.End()

	gctx.Request = gctx.Request.WithContext(ctx)

	h.delete(gctx)
}

func (h *Handler) Fetch(gctx *gin.Context) {
	ctx, span := h.tr.Start(gctx.Request.Context(), "Fetch")
	defer span.End()

	gctx.Request = gctx.Request.WithContext(ctx)

	h.fetch(gctx)
}

func (h *Handler) FetchFile(gctx *gin.Context) {
	ctx, span := h.tr.Start(gctx.Request.Context(), "FetchFile")
	defer span.End()

	gctx.Request = gctx.Request.WithContext(ctx)

	h.fetchFile(gctx)
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
