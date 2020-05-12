package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/wangcheng0509/gpkg/e"
)

type Gin struct {
	C *gin.Context
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code:    errCode,
		Message: e.GetMsg(errCode),
		Data:    data,
	})
	return
}

func (g *Gin) Ok(rsp interface{}) {
	g.C.JSON(http.StatusOK, rsp)
	return
}
