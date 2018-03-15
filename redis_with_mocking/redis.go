package redis_with_mocking

import (
	"log"
)

type RedisConnI interface {
	Do(commandName string, args ...interface{}) (reply interface{}, err error)
}

type connWrapper struct {
	conn RedisConnI
}

var cw = connWrapper{}

func InsertSaddRedis(key, value string) (int64, error) {
	reply, err := cw.conn.Do("SADD", key, value)
	if err != nil {
		log.Printf("error executing command 'SADD %s %s' on redis\n", key, value)
		return 0, err
	}

	return reply.(int64), nil
}
