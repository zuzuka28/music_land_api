package track

import "github.com/gin-gonic/gin"

type Handler struct {
	delete gin.HandlerFunc
	fetch  gin.HandlerFunc
	create gin.HandlerFunc
	search gin.HandlerFunc
}

func NewHandler(s Service) *Handler {
	return &Handler{
		delete: makeDeleteHandler(s),
		fetch:  makeFetchHandler(s),
		create: makeCreateHandler(s),
		search: makeSearchHandler(s),
	}
}

func (h *Handler) Delete(gctx *gin.Context) {
	h.delete(gctx)
}

func (h *Handler) Fetch(gctx *gin.Context) {
	h.fetch(gctx)
}

func (h *Handler) Create(gctx *gin.Context) {
	h.create(gctx)
}

func (h *Handler) Search(gctx *gin.Context) {
	h.search(gctx)
}
