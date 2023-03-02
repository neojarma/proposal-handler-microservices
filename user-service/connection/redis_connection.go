package connection

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

func GetConnectionRedis() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"),
	})

	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		err := rdb.Ping(context.Background()).Err()
		if err == nil {
			log.Println("success connect to redis")
			return rdb, nil
		}

		log.Println("failed to connect to redis, try again in 1 minute")
		time.Sleep(1 * time.Minute)
	}

	log.Println("failed to connect to redis after", maxRetries, "tries")
	return nil, errors.New("failed connection")

}
