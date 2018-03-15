package redis

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

var conn redis.Conn

func InsertSaddRedis(key, value string) (int64, error) {
	reply, err := conn.Do("SADD", key, value)
	if err != nil {
		log.Printf("error executing command 'SADD %s %s' on redis\n", key, value)
		return 0, err
	}

	return reply.(int64), nil
}
