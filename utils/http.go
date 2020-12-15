package utils

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// HTTPGet Get请求
func HTTPGet(out interface{}, url string, header map[string]string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	rsp, err := client.Do(req)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}
	if err := JSONUnmarshal(string(body), out); err != nil {
		return err
	}
	return nil
}

// HTTPPost POST请求
func HTTPPost(out interface{}, url string, header map[string]string, param string) error {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(param))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	for k, v := range header {
		req.Header.Set(k, v)
	}
	rsp, err := client.Do(req)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}
	if err := JSONUnmarshal(string(body), out); err != nil {
		return err
	}
	return nil
}
