package apollo

import (
	"fmt"

	"github.com/shima-park/agollo"
	"github.com/wangcheng0509/gpkg/loghelp"
)

var ApolloConfig agollo.Agollo

type Apollo struct {
	AppName string
	Path    string
	AppId   string
}

// apollo初始化
// path : 192.168.20.12:8080
// appId: someApp
func InitAgollo(apollo Apollo) {
	ApolloConfig, _ = agollo.New(apollo.Path, apollo.AppId, agollo.AutoFetchOnCacheMiss())
	loghelp.Log(fmt.Sprintf("%s Apollo启动成功", apollo.AppName), "", false)
}

func GetSetting(key, ns string) string {
	return ApolloConfig.Get(key, agollo.WithNamespace(ns))
}
