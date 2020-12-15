package exception

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/wangcheng0509/gpkg/e"
	logHelp "github.com/wangcheng0509/gpkg/exceptionless"

	"github.com/gin-gonic/gin"
	timeF "github.com/xinliangnote/go-util/time"
)

// ErrOption 配置
type ErrOption struct {
	AppName         string
	SystemEmailHost string
	SystemEmailPort int
	SystemEmailUser string
	SystemEmailPass string
	ErrorNotifyUser string
	IsLog           bool
	// URL 消息中心地址
	URL string
	// Webhook 钉钉Webhook
	Webhook string `json:"webhook"`
	// Secret 钉钉密钥
	Secret string `json:"secret"`
}

var errSetting = &ErrOption{}

// Init 初始化
func Init(_errSetting *ErrOption) {
	errSetting = _errSetting
}

// 自定义异常标志
var customFlag = "custom err"

// ThrowCustomErr 抛出自定义错误
func ThrowCustomErr(code int, msg string) {
	panic(fmt.Sprintf("%s||%d||%s", customFlag, code, msg))
}

// ExceptionHandle 异常处理
func ExceptionHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqJSON, _ := c.GetRawData()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqJSON))

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

				sendEmail(c, err, reqJSON)
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

func sendEmail(c *gin.Context, err interface{}, reqJSON []byte) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	DebugStack := ""
	for _, v := range strings.Split(string(debug.Stack()), "\n") {
		DebugStack += v + "<br>"
	}
	subject := fmt.Sprintf("【重要错误】%s 项目出错了！", errSetting.AppName)

	body := strings.ReplaceAll(MailTemplate, "{ErrorMsg}", fmt.Sprintf("%s", err))
	body = strings.ReplaceAll(body, "{RequestTime}", timeF.GetCurrentDate())
	body = strings.ReplaceAll(body, "{RequestURL}", c.Request.Method+"  "+c.Request.Host+c.Request.RequestURI)
	body = strings.ReplaceAll(body, "{RequestUA}", c.Request.UserAgent())
	body = strings.ReplaceAll(body, "{RequestBody}", string(reqJSON))
	body = strings.ReplaceAll(body, "{RequestIP}", c.ClientIP())
	body = strings.ReplaceAll(body, "{DebugStack}", DebugStack)

	SendEmailNotice(subject, body)
	SendDingdingNotice(subject, body)

	if errSetting.IsLog {
		msg := fmt.Sprintf(`Application:%s,
		ClassName:%s,
		Message:%s,
		StackTrace:%s,
		CreatedDate:%s`, errSetting.AppName, c.Request.RequestURI, fmt.Sprintf("%s", err), DebugStack, time.Now().Format("2006-01-02 15:04:05"))
		logHelp.Error(msg, true)
	}
}
