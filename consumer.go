package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Consumer")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Panic(err.Error())
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Panic(err.Error())
	}
	defer ch.Close()

	messages, err := ch.Consume(
		"go-queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	// Process the messages
	// Create a channel that will block the execution of main() until it receives a value
	forever := make(chan bool)
	go func() {
		for d := range messages {
			fmt.Printf("Received Message: %s\n", d.Body)
		}
	}()

	fmt.Println("Successfully Connected to our RabbitMQ instance")
	fmt.Println("- waiting for messages")
	<-forever
}
