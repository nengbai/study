package common

import (
	_ "errors"
	"fmt"
	_ "reflect"
	"strconv"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
)

const (
	RedisURL            = "redis:172.28.39.142:6379"
	redisMaxIdle        = 25  //最大空闲连接数
	redisMaxActive      = 100 //最大的激活连接数
	redisIdleTimeoutSec = 240 //最大空闲连接时间
	//RedisPassword       = ""
)

/*
* Redis 连接池
* NewRedisPool 返回redis连接池
* 参数：  redisURL redis服务； Database redisDB；
* 返回： redis 连接
 */
func NewRedisPool(redisURL string, Database int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     redisMaxIdle,
		MaxActive:   redisMaxActive,
		IdleTimeout: redisIdleTimeoutSec * time.Second,
		Dial: func() (redis.Conn, error) {
			URLs := strings.Join([]string{redisURL, strconv.Itoa(Database)}, "/")
			fmt.Println(URLs)
			c, err := redis.DialURL(URLs)
			//c, err := redis.Dial("tcp",redisURL,redis.DialDatabase(5))
			if err != nil {
				return nil, fmt.Errorf("redis connection error: %s", err)
			}
			/*//验证redis密码
			  if _, authErr := c.Do("AUTH", RedisPassword); authErr != nil {
			      return nil, fmt.Errorf("redis auth password error: %s", authErr)
			  }
			*/
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			if err != nil {
				return fmt.Errorf("ping redis error: %s", err)
			}
			return nil
		},
	}
}
