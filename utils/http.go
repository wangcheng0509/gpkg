package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// HTTPGet Get请求
func HTTPGet(out interface{}, url string, header map[string]string) error {
	errmsg := fmt.Sprintf("url:%s;", url)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		errmsg = "发起请求错误;" + errmsg + err.Error()
		return errors.New(errmsg)
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	rsp, err := client.Do(req)
	if err != nil {
		errmsg = "获取返回错误;" + errmsg + err.Error()
		return errors.New(errmsg)
	}
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		errmsg = "读取body错误;" + errmsg + fmt.Sprintf("rspBody:%s;errInfo:%s", body, err.Error())
		return errors.New(errmsg)
	}
	if err := JSONUnmarshal(string(body), out); err != nil {
		errmsg = "解析body错误;" + errmsg + fmt.Sprintf("rspBody:%s;errInfo:%s", body, err.Error())
		return errors.New(errmsg)
	}
	return nil
}

// HTTPPost POST请求
func HTTPPost(out interface{}, url string, header map[string]string, param string) error {
	errmsg := fmt.Sprintf("url:%s;param:%s;", url, param)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(param))
	if err != nil {
		errmsg = "发起请求错误;" + errmsg + err.Error()
		return errors.New(errmsg)
	}
	req.Header.Set("Content-Type", "application/json")
	for k, v := range header {
		req.Header.Set(k, v)
	}
	rsp, err := client.Do(req)
	if err != nil {
		errmsg = "获取返回错误;" + errmsg + err.Error()
		return errors.New(errmsg)
	}
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		errmsg = "读取body错误;" + errmsg + fmt.Sprintf("rspBody:%s;errInfo:%s", body, err.Error())
		return errors.New(errmsg)
	}
	if err := JSONUnmarshal(string(body), out); err != nil {
		errmsg = "解析body错误;" + errmsg + fmt.Sprintf("rspBody:%s;errInfo:%s", body, err.Error())
		return errors.New(errmsg)
	}
	return nil
}
