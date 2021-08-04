package main

import (
	"fmt"
	"strconv"
	"study/RabbitMQ"
)

func main() {
	// rabbitmq := RabbitMQ.NewRabitMQSample("newProduct")
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("PubSubQueueone", "ExchangePubSub")
	// rabbitmqOne := RabbitMQ.NewRabbitMQRouting("RoutingQueueOne","newRouting", "imooc.topic.one")
	// rabbitmqTow := RabbitMQ.NewRabbitMQRouting("RoutingQueueTow","newRouting", "imooc.topic.tow")
	// rabbitmqOne := RabbitMQ.NewRabbitMQTopic("TopicQueueOne", "newTopic", "imooc.topic.one")
	// rabbitmqTow := RabbitMQ.NewRabbitMQTopic("TopicQueueTow", "newTopic", "imooc.topic.tow")
	msg := make([]string, 0)
	for i := 0; i < 1000; i++ {
		if i%2 == 0 {
			message := "Love you " + strconv.Itoa(i)
			msg = append(msg, message)
			//rabbitmqOne.ProduceMQRouting(message)
			//rabbitmqOne.ProduceMQTopic(message)
			//time.Sleep(1 * time.Second)
		} else {
			message := "you quit " + strconv.Itoa(i)
			msg = append(msg, message)
			//rabbitmqTow.ProduceMQRouting(message)
			//rabbitmqTow.ProduceMQTopic(message)
			//time.Sleep(1 * time.Second)
		}
	}

	// rabbitmq.ProduceMQSimple(msg)
	rabbitmq.ProduceMQPubSub(msg)
	rabbitmq.Destory()
	// rabbitmqOne.Destory()
	// rabbitmqTow.Destory()
	fmt.Println("发送成功!")

}
