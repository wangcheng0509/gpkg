module github.com/wangcheng0509/gpkg

go 1.14

require (
	github.com/chenjiandongx/ginprom v0.0.0-20210617023641-6c809602c38a
	github.com/elazarl/goproxy v0.0.0-20230731152917-f99041a5c027 // indirect
	github.com/gin-gonic/gin v1.9.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/golang-jwt/jwt/v4 v4.4.2
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/gorilla/websocket v1.4.2
	github.com/jinzhu/gorm v1.9.16
	github.com/kevholditch/gokong v6.0.0+incompatible
	github.com/ory/dockertest/v3 v3.10.0 // indirect
	github.com/parnurzeal/gorequest v0.2.16 // indirect
	github.com/phayes/freeport v0.0.0-20220201140144-74d24b5ae9f5 // indirect
	github.com/prometheus/client_golang v1.8.0
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/shima-park/agollo v1.2.7
	github.com/streadway/amqp v0.0.0-20190827072141-edfb9018d271
	github.com/swaggo/gin-swagger v1.6.0
	github.com/xinliangnote/go-util v0.0.0-20200323134426-527984dc34bf
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
	gopkg.in/ory-am/dockertest.v3 v3.8.1 // indirect
	moul.io/http2curl v1.0.0 // indirect
)

replace gopkg.in/ory-am/dockertest.v3 v3.8.1 => github.com/ory/dockertest/v3 v3.8.1
