package cache

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v7"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:6379", os.Getenv("REDIS_URL")), // use default Addr
		DB:   0,                                              // use default DB
	})
}
