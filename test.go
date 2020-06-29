package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/wangcheng0509/gpkg/utils"

	"github.com/wangcheng0509/gpkg/kong"
	"github.com/wangcheng0509/gpkg/log"
	"github.com/wangcheng0509/gpkg/mysqlconn"
	"github.com/wangcheng0509/gpkg/rabbitmq"
	"github.com/wangcheng0509/gpkg/try"

	"github.com/wangcheng0509/gpkg/gredis"

	"github.com/gin-gonic/gin"
	"github.com/wangcheng0509/gpkg/exception"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"
	"github.com/wangcheng0509/gpkg/aes"
	"github.com/wangcheng0509/gpkg/apollo"
	"github.com/wangcheng0509/gpkg/app"
	jwttool "github.com/wangcheng0509/gpkg/jwt"
	"github.com/wangcheng0509/gpkg/ws"

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
	rsp := app.Response{}
	rsp.SetData(nil)
	appG.Ok(rsp)
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

func WebSocketTest() {
	// 启动服务之前异步开启websocket
	// go ws.Manager.Start()
	var c *gin.Context
	conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if error != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	client := &ws.Client{ID: "ClientId", Socket: conn, Send: make(chan []byte)}
	ws.Manager.Register <- client

	go client.Read()
	go client.Write()

	// 可以异步根据ClientId下发消息
	// client := &ws.Client{ID: "ClientId", Socket: ws.Manager.ClientConns["ClientId"].Socket, Send: make(chan []byte)}
	// ws.Manager.Send([]byte("test"), client)
}

func RedisTest() {
	redisConn := "192.168.20.96:6379|192.168.20.46:6379|192.168.20.69:6379|192.168.20.96:6389|192.168.20.46:6389|192.168.20.69:6389"
	gredis.Setup(redisConn)
	test := RedisTestStruct{
		Name:  "aaa",
		Value: "1111",
	}

	if err := gredis.RedisConn.Set("sssskey", &test, 60*time.Minute).Err(); err != nil {
		fmt.Println(err)
	}
	cache := RedisTestStruct{}
	gredis.Get("sssskey", &cache)
	fmt.Println(cache)
}

type RedisTestStruct struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// func (s *RedisTestStruct) MarshalBinary() ([]byte, error) {
// 	return json.Marshal(s)
// }

// func (s *RedisTestStruct) UnmarshalBinary(data []byte) error {
// 	return json.Unmarshal(data, s)
// }

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
		Unique_name  string `json:"unique_name"`
		Guid         string `json:"guid"`
		Avatar       string `json:"avatar"`
		DisplayName  string `json:"displayName"`
		LoginName    string `json:"loginName"`
		EmailAddress string `json:"emailAddress"`
		UserType     string `json:"userType"`
		Time         string `json:"time"`
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
	claimsJson, _ := jwttool.Parse("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1bmlxdWVfbmFtZSI6IndhbmdjaGVuZyIsImd1aWQiOiJjMTc2ZmUzOC1jYWQ5LTQ1NDAtODJhZC0yYzg4NWQ2YjNhYzMiLCJhdmF0YXIiOiIiLCJkaXNwbGF5TmFtZSI6IueOi-aIkCIsImxvZ2luTmFtZSI6IndhbmdjaGVuZyIsImVtYWlsQWRkcmVzcyI6IiIsInVzZXJUeXBlIjoiMCIsIlRpbWUiOiIyMDIwMDUwOTExMzczMSIsIm5iZiI6MTU4ODk5NTQ1MSwiZXhwIjoxNjIwNTMxNDUxLCJpYXQiOjE1ODg5OTU0NTEsImlzcyI6Imluc3BpcnkiLCJhdWQiOiJpbnNwaXJ5In0.xsxkBsVHrIr5uIg2NMZu2vsTHjZ-4fwAB3YXFdURbC0")
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

func DistinctTest() {
	var strarr = []string{"1", "1", "2", "3", "2", "3"}
	strarr = utils.DistinctElement(strarr)
	fmt.Println(strarr)
}

func IsExistItemTest() {
	type Stu struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	stus := []Stu{
		Stu{
			Name: "tome",
			Age:  10,
		},
		Stu{
			Name: "jarry",
			Age:  10,
		},
	}
	stu := Stu{
		Name: "jarry",
		Age:  10,
	}
	fmt.Println(utils.IsExistItem(stu, stus))
}

func CombinSqlTest() {
	sql := "select * from test $WHERE"
	var where []string
	where = append(where, "a=1")
	where = append(where, "b=1")
	where = append(where, "c=1")
	sql = utils.CombinSql(sql, where)
	fmt.Println(sql)
}

func main() {
	RedisTest()
}
