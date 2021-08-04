package main

import "study/RabbitMQ"

func main0() {
	// rabbitmq := RabbitMQ.NewRabitMQSample("newProduct")
	// rabbitmq.ConsumeMQSimple("Hello test!")
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("PubSubQueue", "ExchangePubSub")
	rabbitmq.ConsumeMQPubSub()

}
