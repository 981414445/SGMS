package base

import (
	"fmt"
	"testing"
	"time"
)

func TestGetRedis(t *testing.T) {
	redis := GetRedis()
	defer redis.Close()
	pong, err := redis.Ping().Result()
	fmt.Println(pong, err)
	redis.Set("v1", "hello", 10*time.Second)
	fmt.Println(redis.Get("v1").Result())
}
