package gredis

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

var RedisConn *redis.ClusterClient

// Setup Initialize the Redis instance
// redisConn: 127.0.0.1:6379|127.0.0.1:6389
func Setup(redisConn string) {
	cons := strings.Split(redisConn, "|")

	RedisConn = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        cons,
		ReadTimeout:  50 * time.Millisecond,
		WriteTimeout: 50 * time.Millisecond,
	})
	// RedisConn.Do("SET", "gredis-loading", "success")
	fmt.Println("******************************************************************************")
	fmt.Println("********************************redis启动成功**********************************")
	fmt.Println("******************************************************************************")
}

func Get(key string, v interface{}) {
	jsonStr, _ := RedisConn.Get(key).Result()
	json.Unmarshal([]byte(jsonStr), v)
}
