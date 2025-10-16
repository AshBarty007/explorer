package redis

import (
	"context"
	"fmt"
	"testing"
)

func TestNewRedisClient(t *testing.T) {
	NewRedisClient()

	ctx := context.Background()

	err := RedisClient.Set(ctx, "foo", "bar", 60).Err()
	if err != nil {
		panic(err)
	}

	val, err := RedisClient.Get(ctx, "foo").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("foo", val)

	RedisClient.Del(ctx, "foo")
	RedisClient.Close()

}
