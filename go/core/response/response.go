package response

import (
	"errors"
	"net/http"
)

type Responder interface {
	Response() interface{}
}

const (
	HTTPStatusOKCode      = http.StatusOK
	BadRequestErrorCode   = http.StatusBadRequest
	InternalServerErrCode = http.StatusInternalServerError
)

var (
	BadRequest          = errors.New("bad request")
	InternalServerError = errors.New("internal server error")
)

func Generate(data Responder, err error) (int, interface{}) {
	if err != nil {
		if errors.Is(err, BadRequest) {
			return BadRequestErrorCode, err.Error()
		}

		if errors.Is(err, InternalServerError) {
			return InternalServerErrCode, err.Error()
		}
	}

	return HTTPStatusOKCode, data.Response()

}
