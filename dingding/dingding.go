package dingding

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/wangcheng0509/gpkg/utils"
)

// SendDingdingReq 入参
type SendDingdingReq struct {
	Content DingReq `json:"content"`
	Webhook string  `json:"webhook"`
	Secret  string  `json:"secret"`
}

// SendDingdingMsg 发送钉钉预警
func SendDingdingMsg(req SendDingdingReq) error {
	reqParamStr, _ := json.Marshal(req.Content)
	var dingRsp DingRsp
	webhook := getSecret(req.Webhook, req.Secret)
	if err := utils.HTTPPost(&dingRsp, webhook, nil, string(reqParamStr)); err != nil {
		return err
	}
	if dingRsp.Errcode != 0 {
		msg := fmt.Sprintf("发送钉钉预警消息失败,Webhook:%s,请求参数：%s,返回：%v", req.Webhook, string(reqParamStr), dingRsp)
		return errors.New(msg)
	}
	return nil
}

func getSecret(webhook, secret string) string {
	// 加签
	timestamp := strconv.FormatInt(time.Now().Unix()*1000, 10)
	secretStr, _ := sign(timestamp, secret)
	webhook = webhook + "&timestamp=" + timestamp + "&sign=" + secretStr
	return webhook
}

func sign(timestamp string, secret string) (string, error) {
	stringToSign := fmt.Sprintf("%s\n%s", timestamp, secret)
	h := hmac.New(sha256.New, []byte(secret))
	if _, err := io.WriteString(h, stringToSign); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}
