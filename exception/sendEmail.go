package exception

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/wangcheng0509/gpkg/loghelp"
	"github.com/wangcheng0509/gpkg/utils"
	"github.com/xinliangnote/go-util/mail"
)

// SendEmailNotice 发送邮件通知
func SendEmailNotice(subject, body string) {
	defer func() {
		if err := recover(); err != nil {
			loghelp.Error(subject+"发送邮件错误", fmt.Sprintf("%s", err), true)
		}
	}()

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
}

// SendDingdingNotice 发送钉钉预警
func SendDingdingNotice(subject, body string) error {
	templet := "# {AppName} 异常提醒  \n **{time}**  \n **{errMsg}**  \n {errInfo}"
	dingTemple := templet
	dingTemple = strings.ReplaceAll(dingTemple, "{AppName}", subject)
	dingTemple = strings.ReplaceAll(dingTemple, "{time}", time.Now().Format("2006-1-6 15:4:5"))
	dingTemple = strings.ReplaceAll(dingTemple, "{errMsg}", subject)
	dingTemple = strings.ReplaceAll(dingTemple, "{errInfo}", body)
	reqParam := sendDingdingReq{
		DingReq: dingReq{
			Msgtype: "markdown",
			Markdown: markdown{
				Title: subject + " 异常提醒",
				Text:  dingTemple,
			},
		},
		Webhook: errSetting.Webhook,
		Secret:  errSetting.Secret,
	}
	var rspStr string
	reqByte, _ := json.Marshal(&reqParam)
	if err := utils.HTTPPost(&rspStr, errSetting.URL+"/DingDing", nil, string(reqByte)); err != nil {
		loghelp.Error(subject+" 发送钉钉错误", fmt.Sprintf("%s", err), true)
		return err
	}
	return nil
}
