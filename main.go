package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/wangcheng0509/gpkg/dingding"
	"github.com/wangcheng0509/gpkg/exceptionless"
	"github.com/wangcheng0509/gpkg/loghelp"

	"github.com/wangcheng0509/gpkg/utils"

	"github.com/wangcheng0509/gpkg/kong"
	"github.com/wangcheng0509/gpkg/log"
	"github.com/wangcheng0509/gpkg/mysqlconn"
	"github.com/wangcheng0509/gpkg/rabbitmq"
	"github.com/wangcheng0509/gpkg/try"

	"github.com/wangcheng0509/gpkg/gredis"

	"github.com/wangcheng0509/gpkg/exception"

	"github.com/gin-gonic/gin"

	"github.com/wangcheng0509/gpkg/aes"
	"github.com/wangcheng0509/gpkg/apollo"
	"github.com/wangcheng0509/gpkg/app"
	jwttool "github.com/wangcheng0509/gpkg/jwt"
	"github.com/wangcheng0509/gpkg/ws"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/websocket"
)

func main() {
	JwtTest()
}

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
	gredis.SetupCluster(redisConn, "")
	test := RedisTestStruct{
		Name:  "aaa",
		Value: "1111",
	}

	if err := gredis.Cluster.Set("sssskey", &test, 60*time.Minute).Err(); err != nil {
		fmt.Println(err)
	}
	cache := RedisTestStruct{}
	gredis.Cluster.Get("sssskey")
	fmt.Println(cache)

	gredis.Client.Publish("channel-data_sync_aiot_device_attribute", test)
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
	}
	rabbitmq.Init(mqSetting)
	rabbitmq.SendMsg("", "", []byte("msg"))
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
	claimsJson, _ := jwttool.Parse("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1bmlxdWVfbmFtZSI6InVzZXJuYW1lIiwiZ3VpZCI6IjEyMzQ1NiIsImF2YXRhciI6IkF2YXRhciIsImRpc3BsYXlOYW1lIjoi546L5oiQIiwibG9naW5OYW1lIjoid2FuZ2NoZW5nIiwiZW1haWxBZGRyZXNzIjoid2FuZ2NoZW5nQGluc3BpcnkuY24iLCJ1c2VyVHlwZSI6IjAiLCJ0aW1lIjoiMjAyMDA1MDcxNTIzMTMiLCJhdWQiOiJpbnNwaXJ5IiwiZXhwIjoxNjYwMDQyNjAwLCJpc3MiOiJpbnNwaXJ5In0.w13ZFs1kchDjZz7vl6n2CL-lzbpLhlySsrRtXVdI5h8")
	fmt.Println(claimsJson)
	// json转化
	b := []byte(claimsJson)
	userinfo2 := Claims{}
	json.Unmarshal(b, &userinfo2)
	fmt.Println(userinfo2)
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
		{
			Name: "tome",
			Age:  10,
		},
		{
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

func HttpTest() {
	var rsp string
	if err := utils.HTTPPost(&rsp, "https://www.inspiry.cn/", nil, ""); err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp)
}

func Exceptionless() {
	exSetting := exceptionless.ExSetting{
		AppName:       "gpkg",
		APIKey:        "Jj82AlEi5GGlC8tCBdaheydP7kXBDRcJXK9YRX3V",
		URL:           "http://api.t.dev.pay.fun/message",
		ExceptionMail: "wangcheng@pay.media",
	}
	exceptionless.Init(exSetting)
	exceptionless.Log("exceptionless test", true)
	exceptionless.Error("exceptionless test", true)
}

func DingdingMsgTest() {
	webhook := "https://oapi.dingtalk.com/robot/send?access_token=e34ac31df5d217d872e265c836001bb50315239d608d6c0ac9444c8763b454f4"
	secret := "SEC11c5f7696030690ce6183bd632fe498b19b70d8f2e970889c214d5159702a903"
	templet := "# {AppName} 异常提醒  \n **{time}**  \n **{errMsg}**  \n {errInfo}"
	dingTemple := templet
	dingTemple = strings.ReplaceAll(dingTemple, "{AppName}", "aiot.mq")
	dingTemple = strings.ReplaceAll(dingTemple, "{time}", time.Now().Format("2006-1-6 15:4:5"))
	dingTemple = strings.ReplaceAll(dingTemple, "{errMsg}", "异常消息")
	dingTemple = strings.ReplaceAll(dingTemple, "{errInfo}", "异常明细")
	reqParam := dingding.DingReq{
		Msgtype: "markdown",
		Markdown: dingding.Markdown{
			Title: "aiot.mq 异常提醒",
			Text:  dingTemple,
		},
	}
	req := dingding.SendDingdingReq{
		Content: reqParam,
		Webhook: webhook,
		Secret:  secret,
	}
	dingding.SendDingdingMsg(req)
}

func loghelpTest() {
	loghelp.InitLog(loghelp.LogSetting{
		Application: "gpk",
		DingSetting: dingding.SendDingdingReq{
			Webhook: "https://oapi.dingtalk.com/robot/send?access_token=708446152fb5472907b78b35c332fc7be9a437f45f65efb1240b2dfe0eea97f6",
			Secret:  "SECb3e44230a7f63943c80091c0f67cc5bb386694db303808f34587b9fbb0893468",
		},
	})
	loghelp.Log("logtest", "stacktrace", false)
	loghelp.Error("logtest", "stacktrace", true)
}
