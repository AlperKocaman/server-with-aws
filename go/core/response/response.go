package response

import (
	"errors"
	"net/http"
)

type Responder interface {
	Response() interface{}
}

const (
	HTTPStatusOKCode          = http.StatusOK
	BadRequestErrorCode       = http.StatusBadRequest
	NotFoundErrorCode         = http.StatusNotFound
	InternalServerErrCode     = http.StatusInternalServerError
	ServiceUnavailableErrCode = http.StatusServiceUnavailable
)

var (
	BadRequestError         = errors.New("bad request")
	NotFoundError           = errors.New("not found")
	InternalServerError     = errors.New("internal server error")
	ServiceUnavailableError = errors.New("service unavailable")
)

type Generic struct {
	// example: 200
	Code int `json:"code"`
	// example: true
	Success bool `json:"success"`
}

type GenericSuccess struct {
	Generic
	Data interface{} `json:"data"`
}

type GenericError struct {
	Generic
	// example: service unable
	Message string `json:"message"`
}

func Generate(data Responder, err error) (int, interface{}) {
	if err != nil {
		resp := GenericError{}
		if errors.Is(err, BadRequestError) {
			resp.Generic.Code = BadRequestErrorCode
			resp.Message = err.Error()
		} else if errors.Is(err, NotFoundError) {
			resp.Generic.Code = NotFoundErrorCode
			resp.Message = err.Error()
		} else if errors.Is(err, InternalServerError) {
			resp.Generic.Code = InternalServerErrCode
			resp.Message = err.Error()
		} else if errors.Is(err, ServiceUnavailableError) {
			resp.Generic.Code = ServiceUnavailableErrCode
			resp.Message = err.Error()
		}
		return resp.Code, resp
	}

	resp := GenericSuccess{
		Generic: Generic{
			Code:    HTTPStatusOKCode,
			Success: true,
		},
		Data: data.Response(),
	}

	return HTTPStatusOKCode, resp

}
