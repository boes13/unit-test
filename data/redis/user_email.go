package redis

import (
	"github.com/boes13/unit-test/common"
	"github.com/garyburd/redigo/redis"
)

func GetUserEmail(userID int64) (string, error) {
	redisCon := common.GetRedisConnection()
	return redis.String(redisCon.Do("GET", common.GetRedisKeyUser(userID)))
}