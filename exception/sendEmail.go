package exception

import (
	"fmt"
	"runtime/debug"
	"strings"
	"time"

	"github.com/wangcheng0509/gpkg/log"

	"github.com/gin-gonic/gin"
	"github.com/xinliangnote/go-util/mail"
	timeF "github.com/xinliangnote/go-util/time"
)

func sendEmail(c *gin.Context, err interface{}, reqJson []byte) {
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
	body = strings.ReplaceAll(body, "{RequestBody}", string(reqJson))
	body = strings.ReplaceAll(body, "{RequestIP}", c.ClientIP())
	body = strings.ReplaceAll(body, "{DebugStack}", DebugStack)

	options := &mail.Options{
		MailHost: errSetting.SystemEmailHost,
		MailPort: errSetting.SystemEmailPort,
		MailUser: errSetting.SystemEmailUser,
		MailPass: errSetting.SystemEmailPass,
		MailTo:   errSetting.ErrorNotifyUser,
		Subject:  subject,
		Body:     body,
	}

	mail.Send(options)

	if errSetting.IsLog {
		log.Info(log.LogModel{
			Application: errSetting.AppName,
			ClassName:   c.Request.RequestURI,
			Message:     fmt.Sprintf("%s", err),
			StackTrace:  DebugStack,
			Level:       log.Level_Err,
			CreatedDate: time.Now(),
		})
	}
}
