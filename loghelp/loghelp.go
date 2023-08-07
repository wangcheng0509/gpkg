package loghelp

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/wangcheng0509/gpkg/dingding"
)

var logSetting LogSetting

func InitLog(_logSetting LogSetting) {
	logSetting = _logSetting
}

func Log(msg, stackTrace string, sendDing bool) {
	model := LogModel{
		Application: logSetting.Application,
		Message:     msg,
		StackTrace:  stackTrace,
		Level:       1,
		CreatedDate: time.Now(),
	}
	logStr, _ := json.Marshal(model)
	fmt.Println(string(logStr))
	if sendDing {
		logSetting.DingSetting.Content = dingding.DingReq{
			Markdown: dingding.Markdown{
				Title: fmt.Sprintf("%s Log信息", logSetting.Application),
				Text:  fmt.Sprintf("#### %s \n> %s", logSetting.Application, msg),
			},
		}
		dingding.SendDingdingMsg(logSetting.DingSetting)
	}
}

func Error(msg, stackTrace string, sendDing bool) {
	model := LogModel{
		Application: logSetting.Application,
		Message:     msg,
		StackTrace:  stackTrace,
		Level:       2,
		CreatedDate: time.Now(),
	}
	logStr, _ := json.Marshal(model)
	fmt.Println(string(logStr))
	if sendDing {
		logSetting.DingSetting.Content = dingding.DingReq{
			Msgtype: "markdown",
			Markdown: dingding.Markdown{
				Title: fmt.Sprintf("%s 异常信息", logSetting.Application),
				Text:  fmt.Sprintf("#### %s \n> %s", logSetting.Application, msg),
			},
		}
		dingding.SendDingdingMsg(logSetting.DingSetting)
	}
}
