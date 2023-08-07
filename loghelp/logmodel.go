package loghelp

import (
	"github.com/wangcheng0509/gpkg/dingding"
	"time"
)

type LogModel struct {
	Application string    `json:"application"`
	Message     string    `json:"message"`
	StackTrace  string    `json:"stackTrace"`
	Level       int       `json:"level"`
	CreatedDate time.Time `json:"createdDate"`
}

type LogSetting struct {
	DingSetting dingding.SendDingdingReq `json:"dingSetting"`
	Application string                   `json:"application"`
}
