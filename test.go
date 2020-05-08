package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/wangcheng0509/gpkg/kong"
	"github.com/wangcheng0509/gpkg/log"
	"github.com/wangcheng0509/gpkg/mysqlconn"
	"github.com/wangcheng0509/gpkg/rabbitmq"
	"github.com/wangcheng0509/gpkg/try"

	"github.com/wangcheng0509/gpkg/gredis"

	"github.com/gin-gonic/gin"
	"github.com/wangcheng0509/gpkg/e"
	"github.com/wangcheng0509/gpkg/exception"

	"github.com/dgrijalva/jwt-go"
	"github.com/wangcheng0509/gpkg/aes"
	"github.com/wangcheng0509/gpkg/apollo"
	"github.com/wangcheng0509/gpkg/app"
	jwttool "github.com/wangcheng0509/gpkg/jwt"

	"github.com/chenjiandongx/ginprom"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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

func JwtTest() {
	// 初始化jwt参数
	jwttool.Setup("inspiry888888888")
	// 定义model
	type Claims struct {
		Unique_name  string
		Guid         string
		Avatar       string
		DisplayName  string
		LoginName    string
		EmailAddress string
		UserType     string
		Time         string
		jwt.StandardClaims
	}
	// 生成jwt Token
	nowTime := time.Now()
	userinfo := Claims{
		"username",
		"123456",
		"Avatar",
		"王成",
		"wangcheng",
		"wangcheng@inspiry.cn",
		"0",
		"20200507152313",
		jwt.StandardClaims{
			ExpiresAt: (nowTime.Add(365 * time.Hour)).Unix(),
			Issuer:    "inspiry",
			Audience:  "inspiry",
		},
	}
	token, _ := jwttool.GenerateToken(userinfo)
	fmt.Println(token)
	// 解析token
	claimsJson, _ := jwttool.Parse("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVbmlxdWVfbmFtZSI6InVzZXJuYW1lIiwiR3VpZCI6IjEyMzQ1NiIsIkF2YXRhciI6IkF2YXRhciIsIkRpc3BsYXlOYW1lIjoi546L5oiQIiwiTG9naW5OYW1lIjoid2FuZ2NoZW5nIiwiRW1haWxBZGRyZXNzIjoid2FuZ2NoZW5nQGluc3BpcnkuY24iLCJVc2VyVHlwZSI6IjAiLCJUaW1lIjoiMjAyMDA1MDcxNTIzMTMiLCJhdWQiOiJpbnNwaXJ5IiwiZXhwIjoxNTkwMjIzNDE0LCJpc3MiOiJpbnNwaXJ5In0.hKKpBYmJextMGSPipXO_L5B3S9oRim_cw3EIryTdOZE")
	fmt.Println(claimsJson)
	// json转化
	b := []byte(claimsJson)
	userinfo2 := Claims{}
	json.Unmarshal(b, &userinfo2)
	fmt.Println(userinfo2)
}

func PromethesTest() {
	r := gin.New()
	r.Use(ginprom.PromMiddleware(nil))
	r.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))
}

func SwaggerTest() {
	r := gin.New()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// swagger文档更新命令：swag init
}

func main() {
	JwtTest()
}
