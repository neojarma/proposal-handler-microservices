package broker_redis

import (
	"approval-service/models"
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func ListenRedisEvent(conn *redis.Client, chanName string, callback func(payload string)) {
	ctx := context.Background()
	pubsub := conn.Subscribe(ctx, chanName)
	defer pubsub.Close()

	ch := pubsub.Channel()

	log.Println("ready to listen event from", chanName)
	for msg := range ch {
		log.Println("do event from", chanName)
		callback(msg.Payload)
	}
}

func EmitEvent(event models.Event) error {
	return event.Conn.Publish(event.Context, event.Channel, event.Payload).Err()
}
