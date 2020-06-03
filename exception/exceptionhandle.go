package exception

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wangcheng0509/gpkg/e"
)

type ErrOption struct {
	AppName         string
	SystemEmailHost string
	SystemEmailPort int
	SystemEmailUser string
	SystemEmailPass string
	ErrorNotifyUser string
	IsLog           bool
}

var errSetting = &ErrOption{}

func Init(_errSetting *ErrOption) {
	errSetting = _errSetting
}

// 自定义异常标志
var customFlag = "custom err"

// 抛出自定义错误
func ThrowCustomErr(code int, msg string) {
	panic(fmt.Sprintf("%s||%d||%s", customFlag, code, msg))
}

// 异常处理
func ExceptionHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqJson, _ := c.GetRawData()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqJson))

		defer func() {
			if err := recover(); err != nil {
				if errInfo, ok := err.(string); ok {
					if errSlice := strings.Split(errInfo, "||"); len(errSlice) > 2 {
						if errSlice[0] == customFlag {
							c.JSON(http.StatusUnauthorized, gin.H{
								"code":    errSlice[1],
								"message": errSlice[2],
								"data":    nil,
							})

							c.Abort()
							return
						}
					}
				}

				sendEmail(c, err, reqJson)
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":    e.ERROR,
					"message": e.GetMsg(e.ERROR),
					"data":    nil,
				})
				c.Abort()
				return
			}
		}()
		c.Next()
	}
}
