package exception

import (
	"fmt"
	"strings"
	"time"

	"github.com/wangcheng0509/gpkg/dingding"
	"github.com/wangcheng0509/gpkg/loghelp"
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
func SendDingdingNotice(appName, subject, body string) error {
	templet := "# {AppName} 异常提醒  \n **{time}**  \n **{errMsg}**  \n {errInfo}"
	dingTemple := templet
	dingTemple = strings.ReplaceAll(dingTemple, "{AppName}", appName)
	dingTemple = strings.ReplaceAll(dingTemple, "{time}", time.Now().Format("2006-1-6 15:4:5"))
	dingTemple = strings.ReplaceAll(dingTemple, "{errMsg}", subject)
	dingTemple = strings.ReplaceAll(dingTemple, "{errInfo}", body)
	reqParam := dingding.SendDingdingReq{
		Content: dingding.DingReq{
			Msgtype: "markdown",
			Markdown: dingding.Markdown{
				Title: appName + " 异常提醒",
				Text:  dingTemple,
			},
		},
		Webhook: errSetting.Webhook,
		Secret:  errSetting.Secret,
	}

	if err := dingding.SendDingdingMsg(reqParam); err != nil {
		loghelp.Error(appName+" 发送钉钉错误", fmt.Sprintf("%s", err), true)
		return err
	}
	return nil
}
