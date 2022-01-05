package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	connectRabbitMQ, err := amqp.Dial("amqp://guest:guest@localhost:5673/")
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	messages, err := channelRabbitMQ.Consume(
		"QueueService1",    // queue name
		"",              // consumer
		true,             // auto-ack
		false,           // exclusive
		false,            // no local
		false,            // no wait
		nil,                // arguments
	)
	if err != nil {
		log.Println(err)
	}

	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")

	forever := make(chan bool)

	go func() {
		for message := range messages {
			// For example, show received message in a console.
			log.Printf(" > Received message: %s\n", message.Body)

		}
	}()

	<-forever
}
