package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/wangcheng0509/gpkg/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}

func (g *Gin) Ok(data interface{}) {
	g.C.JSON(http.StatusOK, Response{
		Code: e.SUCCESS,
		Msg:  e.GetMsg(e.SUCCESS),
		Data: data,
	})
	return
}
