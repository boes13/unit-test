package common

import "fmt"

const redisKeyUser = "user:%d"

func GetRedisKeyUser(userID int64) string {
	return fmt.Sprintf(redisKeyUser, userID)
}