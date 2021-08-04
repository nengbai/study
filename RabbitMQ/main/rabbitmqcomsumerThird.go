package main

import "study/RabbitMQ"

func main2() {
	// rabbitmq := RabbitMQ.NewRabitMQSample("newProduct")
	// rabbitmq.ConsumeMQSimple("Hello test!")
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("PubSubQueueTow", "ExchangePubSub")
	rabbitmq.ConsumeMQPubSub()

}
