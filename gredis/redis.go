package gredis

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

var Cluster *redis.ClusterClient
var Client *redis.Client

// Setup Initialize the Redis instance
// redisConn: 127.0.0.1:6379|127.0.0.1:6389
func SetupCluster(redisConn, pwd string) {
	cons := strings.Split(redisConn, "|")

	Cluster = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        cons,
		Password:     pwd,
		ReadTimeout:  50 * time.Millisecond,
		WriteTimeout: 50 * time.Millisecond,
	})
	// RedisConn.Do("SET", "gredis-loading", "success")
	fmt.Println("******************************************************************************")
	fmt.Println("********************************redis启动成功**********************************")
	fmt.Println("******************************************************************************")
}

// Setup Initialize the Redis instance
// redisConn: 127.0.0.1:6379|127.0.0.1:6389
func SetupClient(redisConn, pwd string) {
	Client = redis.NewClient(&redis.Options{
		Addr:         redisConn,
		Password:     pwd,
		ReadTimeout:  50 * time.Millisecond,
		WriteTimeout: 50 * time.Millisecond,
	})
	// RedisConn.Do("SET", "gredis-loading", "success")
	fmt.Println("******************************************************************************")
	fmt.Println("********************************redis启动成功**********************************")
	fmt.Println("******************************************************************************")
}
