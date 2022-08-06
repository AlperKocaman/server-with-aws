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
	ServiceUnableErrCode  = http.StatusServiceUnavailable
)

var (
	BadRequest          = errors.New("bad request")
	InternalServerError = errors.New("internal server error")
	ServiceUnableError  = errors.New("service unable")
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
		if errors.Is(err, BadRequest) {
			resp.Generic.Code = BadRequestErrorCode
			resp.Message = err.Error()
		} else if errors.Is(err, InternalServerError) {
			resp.Generic.Code = InternalServerErrCode
			resp.Message = err.Error()
		} else if errors.Is(err, ServiceUnableError) {
			resp.Generic.Code = ServiceUnableErrCode
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
