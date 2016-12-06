package db

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	RedisPool *redis.Pool
)

func InitRedisPool(host string, port string, password string, dbname int64) *redis.Pool {
	RedisPool = &redis.Pool{
		MaxIdle:     3,
		MaxActive:   100,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", host+":"+port)
			if err != nil {
				panic(err.Error())
			}

			//if _, err := con.Do("AUTH", password); err != nil {
			//	con.Close()
			//	return nil, err
			//}

			con.Do("SELECT", dbname)
			return con, err
		},
	}

	return RedisPool
}
