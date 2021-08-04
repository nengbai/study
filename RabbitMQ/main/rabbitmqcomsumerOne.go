package main

import (
	"fmt"
	"study/RabbitMQ"
)

func main1() {
	rabbitmqOne := RabbitMQ.NewRabbitMQTopic("TopicQueueOne", "newTopic", "imooc.*.one")
	rabbitmqOne.ConsumeMQTopic()
	// rabbitmqOne := RabbitMQ.NewRabbitMQRouting("RoutingQueueOne","newRouting", "imooc.topic.one")
	// rabbitmqOne.ConsumeMQRouting()
	rabbitmqOne.Destory()
	fmt.Println("接收成功!")
}
