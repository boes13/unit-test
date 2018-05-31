package common

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

type DataProviderI interface {
	GetRedisConnection() redis.Conn
	GetDBObject() *sql.DB
}

type DataProviderWrapper struct {
	DataProvider DataProviderI
}

type dp struct{}

var dataProviderWrapper DataProviderWrapper


var redisPool redis.Pool
var db *sql.DB

func InitDataConnection() error {
	dataProviderWrapper = DataProviderWrapper{dp{}}

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

func (dp) GetRedisConnection() redis.Conn {
	return redisPool.Get()
}

func (dp) GetDBObject() *sql.DB {
	return db
}

func GetDataProvider() DataProviderWrapper {
	return dataProviderWrapper
}

func SetDataProvider(i DataProviderI) {
	dataProviderWrapper.DataProvider = i
}
