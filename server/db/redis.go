package db

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	RedisPool *redis.Pool
)

func InitPool(dataSourceName string, dbname int64, password string) *redis.Pool {
	RedisPool = &redis.Pool{
		MaxIdle:     3,
		MaxActive:   100,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", dataSourceName)
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
