package album

import "github.com/gin-gonic/gin"

type Handler struct {
	tr Tracer

	delete gin.HandlerFunc
	fetch  gin.HandlerFunc
	create gin.HandlerFunc
	update gin.HandlerFunc
}

func NewHandler(s Service, tr Tracer) *Handler {
	checkaccess := albumActionAccessControl(s)

	return &Handler{
		tr:     tr,
		delete: checkaccess(makeDeleteHandler(s)),
		fetch:  checkaccess(makeFetchHandler(s)),
		create: makeCreateHandler(s),
		update: checkaccess(makeUpdateHandler(s)),
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

func (h *Handler) Create(gctx *gin.Context) {
	ctx, span := h.tr.Start(gctx.Request.Context(), "Create")
	defer span.End()

	gctx.Request = gctx.Request.WithContext(ctx)

	h.create(gctx)
}

func (h *Handler) Update(gctx *gin.Context) {
	ctx, span := h.tr.Start(gctx.Request.Context(), "Update")
	defer span.End()

	gctx.Request = gctx.Request.WithContext(ctx)

	h.update(gctx)
}
