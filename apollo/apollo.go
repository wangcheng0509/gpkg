package apollo

import (
	"fmt"

	"github.com/shima-park/agollo"
)

var ApolloConfig agollo.Agollo

// apollo初始化
// path : 192.168.20.12:8080
// appId: someApp
func InitAgollo(path, appId string) {
	ApolloConfig, _ = agollo.New(path, appId, agollo.AutoFetchOnCacheMiss())
	fmt.Println("**************************************************************")
	fmt.Println("****************Apollo启动成功*********************************")
	fmt.Println("****************" + ApolloConfig.Get("AppName", agollo.WithNamespace("Uei")) + "*******************************")
	fmt.Println("**************************************************************")
}

func GetSetting(key, ns string) string {
	return ApolloConfig.Get(key, agollo.WithNamespace(ns))
}
