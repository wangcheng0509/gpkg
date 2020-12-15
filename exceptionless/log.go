package exceptionless

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/wangcheng0509/gpkg/utils"
)

// ExSetting 配置
type ExSetting struct {
	AppName       string
	APIKey        string
	URL           string
	ExceptionMail string
	Webhook       string `json:"webhook"`
	Secret        string `json:"secret"`
}

var exSetting ExSetting

// Init 初始化
func Init(_exSetting ExSetting) {
	exSetting = _exSetting
}

// Log 记录日志
func Log(msg string, isEmail bool) error {
	req := exLessReq{
		Type:    1,
		APIKey:  exSetting.APIKey,
		Message: msg,
	}
	var rspStr string
	reqByte, _ := json.Marshal(&req)
	if err := utils.HTTPPost(&rspStr, exSetting.URL+"/Log", nil, string(reqByte)); err != nil {
		fmt.Println("请求写入日志错误：")
		fmt.Println(err)
		return err
	}
	var rsp exLessRsp
	if err := json.Unmarshal([]byte(rspStr), &rsp); err != nil {
		fmt.Println("解析写入日志返回错误：")
		fmt.Println(err)
		fmt.Println(exSetting.URL + "/Log")
		fmt.Println(string(reqByte))
		return err
	}
	if rsp.Code != 200 {
		return errors.New(rsp.Message)
	}
	if isEmail {
		sendEmailNotice(exSetting.AppName+" Log", msg)
		sendDingdingNotice(exSetting.AppName+" Error", msg)
		fmt.Println("发送邮件成功")
	}
	return nil
}

// Error 记录异常
func Error(msg string, isEmail bool) error {
	req := exLessReq{
		Type:    2,
		APIKey:  exSetting.APIKey,
		Message: msg,
	}
	var rspStr string
	reqByte, _ := json.Marshal(&req)
	if err := utils.HTTPPost(&rspStr, exSetting.URL+"/Log", nil, string(reqByte)); err != nil {
		fmt.Println("请求写入日志错误：")
		fmt.Println(err)
		return err
	}
	var rsp exLessRsp
	if err := json.Unmarshal([]byte(rspStr), &rsp); err != nil {
		fmt.Println("解析写入日志返回错误：")
		fmt.Println(err)
		fmt.Println(exSetting.URL + "/Log")
		fmt.Println(string(reqByte))
		return err
	}
	if rsp.Code != 200 {
		fmt.Println("日志返回错误：")
		fmt.Println(exSetting.URL + "/Log")
		fmt.Println(string(reqByte))
		fmt.Println(rsp.Message)
		return errors.New(rsp.Message)
	}
	if isEmail {
		sendEmailNotice(exSetting.AppName+" Error", msg)
		sendDingdingNotice(exSetting.AppName+" Error", msg)
		fmt.Println("发送邮件成功")
	}
	return nil
}

// EmailReq 邮件配置
type EmailReq struct {
	Channel      string `json:"Channel"`
	Subject      string `json:"Subject"`
	Content      string `json:"Content"`
	ToCollection string `json:"ToCollection"`
}

func sendEmailNotice(subject, body string) error {
	req := EmailReq{
		Subject:      subject,
		Content:      utils.BaseEncode(body),
		ToCollection: exSetting.ExceptionMail,
	}
	var rspStr string
	reqByte, _ := json.Marshal(&req)
	if err := utils.HTTPPost(&rspStr, exSetting.URL+"/Email", nil, string(reqByte)); err != nil {
		fmt.Println("发送邮件错误")
		return err
	}
	return nil
}

func sendDingdingNotice(subject, body string) error {
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
		Webhook: exSetting.Webhook,
		Secret:  exSetting.Secret,
	}
	var rspStr string
	reqByte, _ := json.Marshal(&reqParam)
	if err := utils.HTTPPost(&rspStr, exSetting.URL+"/DingDing", nil, string(reqByte)); err != nil {
		fmt.Println("发送钉钉错误")
		return err
	}
	return nil
}
