package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/miniwaveme/api/src/conf"
)

var redisCo = newPool()

func newPool() *redis.Pool {

	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {

			c, err := redis.Dial(c.GetString("redis_network"), getUrl())
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

func getUrl() string {

	return fmt.Sprintf("%s:%s", c.GetString("redis_url"), c.GetString("redis_port"))
}
