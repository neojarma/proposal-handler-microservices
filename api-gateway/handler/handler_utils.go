package handler

import (
	"api-gateway/models"
	broker "api-gateway/redis"
	"context"
	"encoding/json"
)

// need to pass pointer to model parameter
func (handler *APIHandlerImpl) emitEventToService(ctx context.Context, model any, chanName string) error {
	payload, err := json.Marshal(model)
	if err != nil {
		return err
	}

	event := models.Event{
		Context: ctx,
		Channel: chanName,
		Payload: payload,
		Conn:    handler.RedisConn,
	}

	return broker.EmitEvent(event)
}
