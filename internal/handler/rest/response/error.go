package response

import (
	"errors"
	"net/http"

	"github.com/zuzuka28/music_land_api/internal/model"
)

type Error struct {
	Message string `json:"message"`
}

func NewError(e error) (code int, err Error) {
	code = http.StatusInternalServerError
	msg := http.StatusText(code)

	if errors.Is(e, model.ErrNotFound) {
		code = http.StatusNotFound
		msg = e.Error()
	}

	if errors.Is(e, model.ErrNotValid) {
		code = http.StatusBadRequest
		msg = e.Error()
	}

	return code, Error{
		Message: msg,
	}
}
