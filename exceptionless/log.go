package exceptionless

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/wangcheng0509/gpkg/utils"
)

type ExSetting struct {
	AppName       string
	ApiKey        string
	Url           string
	ExceptionMail string
}

var exSetting ExSetting

func Init(_exSetting ExSetting) {
	exSetting = _exSetting
}

type exLessReq struct {
	Type    int    `json:"type"`    // 1 log，2 error
	ApiKey  string `json:"apiKey"`  // apiKey
	Message string `json:"message"` // message
}

type exLessRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Log(msg string, isEmail bool) error {
	req := exLessReq{
		Type:    1,
		ApiKey:  exSetting.ApiKey,
		Message: msg,
	}
	var rspStr string
	reqByte, _ := json.Marshal(&req)
	if err := utils.HttpPost(&rspStr, exSetting.Url+"/Log", nil, string(reqByte)); err != nil {
		fmt.Println("请求写入日志错误：")
		fmt.Println(err)
		return err
	}
	var rsp exLessRsp
	if err := json.Unmarshal([]byte(rspStr), &rsp); err != nil {
		fmt.Println("解析写入日志返回错误：")
		fmt.Println(err)
		fmt.Println(exSetting.Url + "/Log")
		fmt.Println(string(reqByte))
		return err
	}
	if rsp.Code != 200 {
		return errors.New(rsp.Message)
	}
	if isEmail {
		sendEmailNotice(exSetting.AppName+" Log", msg)
		fmt.Println("发送邮件成功")
	}
	return nil
}

func Error(msg string, isEmail bool) error {
	req := exLessReq{
		Type:    2,
		ApiKey:  exSetting.ApiKey,
		Message: msg,
	}
	var rspStr string
	reqByte, _ := json.Marshal(&req)
	if err := utils.HttpPost(&rspStr, exSetting.Url+"/Log", nil, string(reqByte)); err != nil {
		fmt.Println("请求写入日志错误：")
		fmt.Println(err)
		return err
	}
	var rsp exLessRsp
	if err := json.Unmarshal([]byte(rspStr), &rsp); err != nil {
		fmt.Println("解析写入日志返回错误：")
		fmt.Println(err)
		fmt.Println(exSetting.Url + "/Log")
		fmt.Println(string(reqByte))
		return err
	}
	if rsp.Code != 200 {
		fmt.Println("日志返回错误：")
		fmt.Println(exSetting.Url + "/Log")
		fmt.Println(string(reqByte))
		fmt.Println(rsp.Message)
		return errors.New(rsp.Message)
	}
	if isEmail {
		sendEmailNotice(exSetting.AppName+" Error", msg)
		fmt.Println("发送邮件成功")
	}
	return nil
}

type EmailReq struct {
	Channel      string `json:"Channel"`
	Subject      string `json:"Subject"`
	Content      string `json:"Content"`
	ToCollection string `json:"ToCollection"`
}

func sendEmailNotice(subject, body string) error {
	req := EmailReq{
		Subject:      subject,
		Content:      body,
		ToCollection: exSetting.ExceptionMail,
	}
	var rspStr string
	reqByte, _ := json.Marshal(&req)
	if err := utils.HttpPost(&rspStr, exSetting.Url+"/Email", nil, string(reqByte)); err != nil {
		fmt.Println("发送邮件错误")
		return err
	}
	return nil
}
