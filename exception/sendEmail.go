package exception

import (
	"fmt"

	"github.com/xinliangnote/go-util/mail"
)

func SendEmailNotice(subject, body string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
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
