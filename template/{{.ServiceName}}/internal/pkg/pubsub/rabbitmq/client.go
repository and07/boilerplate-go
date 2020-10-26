package rabbitmq

import (
	"context"
	"log"

	"github.com/{{.User}}/{{.ServiceName}}/internal/pkg/pubsub"
	"github.com/opentracing/opentracing-go"
	"github.com/streadway/amqp"
)

// NewRabbitmqClient ...
func NewRabbitmqClient(ctx context.Context, connectionString string) (*amqp.Connection, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "NewRabbitmqClient")
	defer span.Finish()
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// New ...
func New(ctx context.Context, connectionString string) (pubsub.PubSuber, error) {
	conn, err := NewRabbitmqClient(ctx, connectionString)
	if err != nil {
		return nil, err
	}
	chSub, errChannel := conn.Channel() // Get a channel from the connection
	if errChannel != nil {
		return nil, errChannel
	}
	chPub, errChannel := conn.Channel() // Get a channel from the connection
	if errChannel != nil {
		return nil, errChannel
	}
	return &clientRabbit{conn: conn, chSub: chSub, chPub: chPub}, err
}

type clientRabbit struct {
	conn  *amqp.Connection
	chSub *amqp.Channel
	chPub *amqp.Channel
}

// Publish ...
func (m *clientRabbit) Publish(ctx context.Context, body []byte, queueName string) error {
	if m.conn == nil {
		panic("Tried to send message before connection was initialized. Don't do that.")
	}

	queue, errQueueDeclare := m.chPub.QueueDeclare( // Declare a queue that will be created if not exists with some args
		queueName, // our queue name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		amqp.Table{
			"x-dead-letter-exchange":    "",
			"x-dead-letter-routing-key": queueName + ".XQ",
		}, // arguments
	)
	if errQueueDeclare != nil {
		return errQueueDeclare
	}

	errQos := m.chPub.Qos(20, 0, false)
	if errQos != nil {
		return errQos
	}

	// Publishes a message onto the queue.
	errPublish := m.chPub.Publish(
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			Headers: amqp.Table{},
			//ContentType:  "application/json",
			Body:         body, // Our JSON body as []byte
			DeliveryMode: amqp.Persistent,
			Priority:     0,
		})
	//fmt.Printf("A message was sent to queue %v: %s", queueName, body)
	return errPublish
}

// Subscribe ...
func (m *clientRabbit) Subscribe(ctx context.Context, queueName string, consumerName string, handlerFunc func(amqp.Delivery)) error {
	log.Printf("Declaring Queue (%s)", queueName)
	queue, errQueueDeclare := m.chSub.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when usused
		false,     // exclusive
		false,     // noWait
		amqp.Table{
			"x-dead-letter-exchange":    "",
			"x-dead-letter-routing-key": queueName + ".XQ",
		}, // arguments
	)
	if errQueueDeclare != nil {
		return errQueueDeclare
	}

	errQos := m.chSub.Qos(20, 0, false)
	if errQos != nil {
		return errQos
	}

	msgs, errConsume := m.chSub.Consume(
		queue.Name,   // queue
		consumerName, // consumer
		false,        // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if errConsume != nil {
		return errConsume
	}

	go consumeLoop(msgs, handlerFunc)
	return nil
}

func (m *clientRabbit) Close() {
	if m.chSub != nil {
		m.chSub.Close()
	}
	if m.chPub != nil {
		m.chPub.Close()
	}
	if m.conn != nil {
		m.conn.Close()
	}
}

func consumeLoop(deliveries <-chan amqp.Delivery, handlerFunc func(d amqp.Delivery)) {
	for d := range deliveries {
		// Invoke the handlerFunc func we passed as parameter.
		handlerFunc(d)
	}
}
