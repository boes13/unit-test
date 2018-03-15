package redis

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

func Init() error {
	var err error
	conn, err = redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Println("error dialing redis:", err.Error())
	}
	return err
}