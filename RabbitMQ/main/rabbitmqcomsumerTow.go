package main

import (
	"fmt"
	"study/RabbitMQ"
)

func main3() {

	rabbitmqTow := RabbitMQ.NewRabbitMQTopic("TopicQueueTow", "newTopic", "imooc.*.tow")
	rabbitmqTow.ConsumeMQTopic()
	// rabbitmqTow := RabbitMQ.NewRabbitMQRouting("RoutingQueueTow","newRouting", "imooc.topic.tow")
	// rabbitmqTow.ConsumeMQRouting()
	rabbitmqTow.Destory()
	fmt.Println("ζ₯ζΆζε!")
}
