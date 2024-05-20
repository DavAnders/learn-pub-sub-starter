package pubsub

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func DeclareAndBind(
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	simpleQueueType int,
	) (*amqp.Channel, amqp.Queue, error) {
		const (
			DurableQueueType   = 1
			TransientQueueType = 2
		)

		fmt.Println("Opening channel...")
		ch, err := conn.Channel()
		if err != nil {
			return nil, amqp.Queue{}, err
		}

		durable := simpleQueueType == DurableQueueType
		autoDelete := simpleQueueType == TransientQueueType
		exclusive := simpleQueueType == TransientQueueType

		fmt.Printf("Declaring queue: %s\n", queueName)
		queue, err := ch.QueueDeclare(
			queueName,
			durable,
			autoDelete,
			exclusive,
			false, // noWait
			nil, // args
		)
		if err != nil {
			return nil, amqp.Queue{}, err
		}

		err = ch.QueueBind(queue.Name, key, exchange, false, nil)
		if err != nil {
			return nil, amqp.Queue{}, err
		}

		fmt.Println("Queue declared and bound successfully.")
		return ch, queue, nil
}