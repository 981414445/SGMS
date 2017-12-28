package base

import (
	"SGMS/domain/config"

	"gopkg.in/redis.v5"
)

func GetRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Redis,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
