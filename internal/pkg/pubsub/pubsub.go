package pubsub

import (
	"context"

	"github.com/streadway/amqp"
)

// PubSuber ...
type PubSuber interface {
	//ConnectToBroker(connectionString string)
	Publish(ctx context.Context, msg []byte, exchangeName string, exchangeType string) error
	//PublishOnQueue(msg []byte, queueName string) error
	Subscribe(ctx context.Context, exchangeName string, exchangeType string, consumerName string, handlerFunc func(amqp.Delivery)) error
	//SubscribeToQueue(queueName string, consumerName string, handlerFunc func(amqp.Delivery)) error
	Close()
}
