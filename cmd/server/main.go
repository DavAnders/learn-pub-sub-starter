package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril server...")
	
	const connStr = "amqp://guest:guest@localhost:5672/"
	
	conn, err := amqp.Dial(connStr)
	if err != nil {
		fmt.Printf("could not connect to RabbitMQ: %v\n", err)
		return
	}
	defer conn.Close()

	newChannel, err := conn.Channel()
	if err != nil {
		fmt.Printf("could not open channel: %v\n", err)
		return
	}
	
	err = pubsub.PublishJSON(newChannel, routing.ExchangePerilDirect, routing.PauseKey, routing.PlayingState{IsPaused: true})
	if err != nil {
		fmt.Printf("could not publish JSON: %v\n", err)
		return
	}

	fmt.Println("connected to RabbitMQ")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	// Wait for signal
	<-signalChan

	fmt.Println("Received signal, shutting down...")
}
