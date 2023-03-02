package broker

import (
	model "api-gateway/models"
)

func EmitEvent(event model.Event) error {
	return event.Conn.Publish(event.Context, event.Channel, event.Payload).Err()
}
