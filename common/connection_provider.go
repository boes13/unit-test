package common

import (
	"database/sql"
	"time"

	"github.com/garyburd/redigo/redis"
	"fmt"
)

var redisPool redis.Pool
var db *sql.DB

func InitDataConnection() error {
	// configuration from files
	redisTimeout := 60
	redisMaxConn := 30
	redisNetwork := "tcp"
	redisAddress := "localhost:6379"
	dbUser := "boes"
	dbPassword := "boes"
	dbName := "demo"
	dbDriver := "mysql"

	redisPool = redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial(redisNetwork, redisAddress)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		IdleTimeout: time.Duration(redisTimeout) * time.Second,
		MaxActive:   redisMaxConn,
		MaxIdle:     redisMaxConn / 2,
		Wait:        true,
	}

	var err error
	db, err = sql.Open(dbDriver, fmt.Sprintf("%s:%s@/%s", dbUser, dbPassword, dbName))
	return err
}

func GetRedisConnection() redis.Conn {
	return redisPool.Get()
}

func GetDBObject() *sql.DB {
	return db
}
