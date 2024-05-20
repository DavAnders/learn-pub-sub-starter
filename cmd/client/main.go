package main

import (
	"fmt"
	"time"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/gamelogic"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril client...")

	const connStr = "amqp://guest:guest@localhost:5672/"

	conn, err := amqp.Dial(connStr)
	if err != nil {
		fmt.Printf("could not connect to RabbitMQ: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Println("connected to RabbitMQ")

	username, err := gamelogic.ClientWelcome()
	if err != nil {
		fmt.Printf("could not welcome client: %v\n", err)
		return
	}

	queueName := fmt.Sprintf("%s.%s", routing.PauseKey, username)

	channel, _, err := pubsub.DeclareAndBind(conn, "peril_direct", queueName, routing.PauseKey, 2) // Transient queue
	if err != nil {
		fmt.Printf("could not declare and bind: %v\n", err)
		return
	}
	defer channel.Close()

	fmt.Println("Press Ctrl+C to exit")
	for {
		time.Sleep(1 * time.Second)
	}
}
