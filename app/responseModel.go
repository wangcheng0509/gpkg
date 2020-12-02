package app

import (
	"github.com/wangcheng0509/gpkg/e"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *Response) SetSuccess(message string) {
	r.Message = message
	r.Code = e.SUCCESS
}

func (r *Response) SetFailed(message string) {
	r.Message = message
	r.Code = e.FAILED
}

func (r *Response) SetError(message string) {
	r.Message = message
	r.Code = e.ERROR
}

func (r *Response) SetData(data interface{}) {
	r.Message = e.GetMsg(e.SUCCESS)
	r.Code = e.SUCCESS
	r.Data = data
}
