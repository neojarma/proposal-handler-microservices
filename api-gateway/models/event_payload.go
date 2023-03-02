package models

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Event struct {
	Context context.Context
	Channel string
	Payload []byte
	Conn    *redis.Client
}
