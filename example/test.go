package example

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/wangcheng0509/gpkg/kong"
	"github.com/wangcheng0509/gpkg/log"
	"github.com/wangcheng0509/gpkg/mysqlconn"
	"github.com/wangcheng0509/gpkg/rabbitmq"
	"github.com/wangcheng0509/gpkg/try"

	"github.com/wangcheng0509/gpkg/gredis"

	"github.com/gin-gonic/gin"
	"github.com/wangcheng0509/gpkg/e"
	"github.com/wangcheng0509/gpkg/exception"

	"github.com/wangcheng0509/gpkg/aes"
	"github.com/wangcheng0509/gpkg/apollo"
	"github.com/wangcheng0509/gpkg/app"
)

func AesTest() {
	appKey := "1234567890123456"
	base64Str, _ := base64.StdEncoding.DecodeString("aestest")
	reqStr, err := aes.AesECBDecrypt(base64Str, []byte(appKey))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reqStr)
}

func ApolloTest() {
	setting := apollo.Apollo{
		Path:  "127.0.0.1:8080",
		AppId: "AppId",
	}
	apollo.InitAgollo(setting)
	apollo.GetSetting("key", "namespace")
}

func AppRspTest(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, "success")
}

func ExceptionMiddlewareTest() {
	exceptionSetting := exception.ErrOption{
		SystemEmailHost: "smtp.exmail.qq.com",
		SystemEmailPort: 587,
		SystemEmailUser: "tmmp@qq.cn",
		SystemEmailPass: "123456",
		ErrorNotifyUser: "tmmp@qq.cn",
		IsLog:           true,
	}
	exception.Init(&exceptionSetting)
	r := gin.New()
	r.Use(exception.ExceptionHandle())
}

func RedisTest() {
	redisConn := "127.0.0.1:6379|127.0.0.2:6379"
	gredis.Setup(redisConn)
	gredis.RedisConn.Set("key", "value", 60*1000)
	cache, _ := gredis.RedisConn.Get("key").Result()
	fmt.Println(cache)
}

func KongTest() {
	routeprotocol := strings.Split("http,https", ",")
	routehost := strings.Split("域名1,域名2", ",")
	kongSetting := kong.Kong{
		KongHost:        "http://kong-proxy:8001",
		UpStreamName:    "UpStreamName",
		TargetPath:      "节点名",
		TargetPort:      "节点端口",
		TargetWeight:    100,
		ServiceName:     "ServiceName",
		ServiceProtocol: "http",
		ServicePort:     80,
		RouteProtocol:   routeprotocol,
		RouteHost:       routehost,
	}
	kong.InitKong(kongSetting)
}

func LogTest() {
	logDBSetting := mysqlconn.Database{
		Type:     "mysql",
		User:     "User",
		Password: "Password",
		Host:     "mysql-0.mysql:3306",
		DBName:   "DBName",
	}
	log.InitDBLog(logDBSetting)
}

func RabbitTest() {
	mqSetting := rabbitmq.Mq{
		Host:     "rabbitmq-Host",
		Port:     "5672",
		Username: "username",
		Pwd:      "pwd",
		Vh:       "Virtual Hosts",
		Queue:    "Queue",
	}
	rabbitmq.Init(mqSetting)
	rabbitmq.SendMsg([]byte("msg"))
}

func TryTest() error {
	var err interface{}
	try.Try(func() {
		defer func() {
			if err = recover(); err != nil {
				try.Throw(1, err.(string))
			}
		}()
		panic("test")
	}).Catch(1, func(e try.Exception) {
		err = errors.New(e.Msg)
	}).Catch(2, func(e try.Exception) {
		err = errors.New(e.Msg)
	}).Finally(func() {})

	return err.(error)
}
