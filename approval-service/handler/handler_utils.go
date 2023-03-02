package handler

import (
	model "approval-service/models"
	broker "approval-service/redis"
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
)

// emit event to update building status
func EmitEventToBuildingService(req any, redisConn *redis.Client, chanName string) {

	bytes, err := json.Marshal(req)
	if err != nil {
		log.Println("error while marshalling payload")
	}

	event := model.Event{
		Context: context.Background(),
		Channel: chanName,
		Payload: bytes,
		Conn:    redisConn,
	}

	broker.EmitEvent(event)
}
