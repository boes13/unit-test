package main

import (
	"log"

	redis "github.com/boes13/unit-test/redis"
	//redis "github.com/boes13/unit-test/redis_with_mocking"
)

func main() {
	err := redis.Init()
	if err != nil {
		log.Println("error initializing module:", err.Error())
		return
	}

	key := "SomeKey123"
	value := "123456"
	count, err := redis.InsertSaddRedis(key, value)
	if err != nil {
		log.Println("error calling function InsertSaddRedis:", err.Error())
		return
	}
	log.Println("output function:", count)
}
