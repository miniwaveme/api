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

			c, err := redis.Dial(conf.C().GetString("redis_network"), getUrl())
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}

}

func Do(commandName string, args ...interface{}) (reply interface{}, err error) {

	c := redisCo.Get()
	defer c.Close()

	return c.Do(commandName, args)
}

func getUrl() string {

	return fmt.Sprintf("%s:%s", conf.C().GetString("redis_url"), conf.C().GetString("redis_port"))
}
