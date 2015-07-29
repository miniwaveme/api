package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/miniwaveme/api/src/config"
)

var (
	gRedisConn = newPool()
)

func newPool() *redis.Pool {

	appConfig := config.GetConfig()

	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(appConfig.GetString("redis_network"), fmt.Sprintf("%s:%s", appConfig.GetString("redis_url"), appConfig.GetString("redis_port")))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}

}

func Do(commandName string, args ...interface{}) (reply interface{}, err error) {

	c := gRedisConn.Get()
	defer c.Close()

	return c.Do(commandName, args)
}
