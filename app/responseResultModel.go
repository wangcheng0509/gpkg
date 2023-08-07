package app

import "github.com/wangcheng0509/gpkg/e"

type ResponseResultModel struct {
	Response
	TotalCount int `json:"totalCount"`
}

func (r *ResponseResultModel) SetData(data interface{}, totalCount int) {
	r.Message = e.GetMsg(e.SUCCESS)
	r.Code = e.SUCCESS
	r.Data = data
	r.TotalCount = totalCount
}
