package app

import "github.com/wangcheng0509/gpkg/e"

type ResponseResultModel struct {
	Response
	TotalCount int64
}

func (r ResponseResultModel) SetData(data interface{}, totalCount int64) {
	r.Message = e.GetMsg(e.SUCCESS)
	r.Code = e.SUCCESS
	r.Data = data
	r.TotalCount = totalCount
}
