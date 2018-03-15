package redis_with_mocking

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

type RedisI interface {
	Dial(network, address string, options ...redis.DialOption) (redis.Conn, error)
}

type redisWrapper struct {
	redis RedisI
}

type redisImpl struct{}

var rdsWrapper = redisWrapper{redisImpl{}}
var network = "tcp"
var address = "localhost:6379"

func Init() error {
	conn, err := rdsWrapper.redis.Dial(network, address)
	if err != nil {
		log.Println("error dialing redis:", err.Error())
	}
	cw.conn = conn
	return err
}

func (r redisImpl) Dial(network, address string, options ...redis.DialOption) (redis.Conn, error) {
	return redis.Dial(network, address, options...)
}
