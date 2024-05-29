package redis

import (
	"encoding/json"
	"errors"
)

type MessagePublisher struct {
	client Redis
}

func NewMessagePublisher(client Redis) *MessagePublisher {
	return &MessagePublisher{
		client,
	}
}

func (p *MessagePublisher) PublishMessages(message interface{}, queueName string) error {
	serializedMessage, err := json.Marshal(message)
	if err != nil {
		return errors.New("failed to serialize message: " + err.Error())
	}

	err = p.client.RedisClient.Publish(queueName, serializedMessage).Err()
	if err != nil {
		return errors.New("failed to publish message: " + err.Error())
	}

	return nil
}
