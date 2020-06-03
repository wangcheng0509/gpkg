package apollo

import (
	"fmt"

	"github.com/shima-park/agollo"
)

var ApolloConfig agollo.Agollo

type Apollo struct {
	Path  string
	AppId string
}

// apollo初始化
// path : 192.168.20.12:8080
// appId: someApp
func InitAgollo(apollo Apollo) {
	ApolloConfig, _ = agollo.New(apollo.Path, apollo.AppId, agollo.AutoFetchOnCacheMiss())
	fmt.Println("**************************************************************")
	fmt.Println("****************Apollo启动成功*********************************")
	fmt.Println("****************" + ApolloConfig.Get("AppName", agollo.WithNamespace("Uei")) + "*******************************")
	fmt.Println("**************************************************************")
}

func GetSetting(key, ns string) string {
	return ApolloConfig.Get(key, agollo.WithNamespace(ns))
}
