package RabbitMQ

import (
	"fmt"
	"log"
	"study/common"

	"github.com/streadway/amqp"
)

const MQURL = "amqp://guest:guest@172.26.107.180:5672/"

type RabbitMQ struct {
	conn         *amqp.Connection
	channelname  *amqp.Channel
	Queuename    string
	Exchangename string
	Routekey     string
	Mqurl        string
}

func (mq *RabbitMQ) FailOnError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}

}

func NewRabbitMQ(queuename, exchangename, routekey string) *RabbitMQ {
	rabbitmq := &RabbitMQ{Queuename: queuename, Exchangename: exchangename, Routekey: routekey}
	var err error
	rabbitmq.Mqurl = MQURL
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	// defer rabbitmq.conn.Close()
	rabbitmq.FailOnError(err, "Fail to Connect RabbitMQ Server error")
	//println(rabbitmq.conn.Channel())
	rabbitmq.channelname, err = rabbitmq.conn.Channel()
	// defer rabbitmq.channelname.Close()
	rabbitmq.FailOnError(err, "Fail to connect RabbitMQ channel error")
	return rabbitmq
}

func (mq *RabbitMQ) Destory() {
	mq.channelname.Close()
	mq.conn.Close()
}

// Create MQ intance with simple
func NewRabitMQSample(queuename string) *RabbitMQ {
	rabbitmq := NewRabbitMQ(queuename, "", "")
	return rabbitmq

}

// Create MQ instance with publish/subsribe mode
func NewRabbitMQPubSub(queuename, exchangename string) *RabbitMQ {
	rabbitmq := NewRabbitMQ(queuename, exchangename, "")
	return rabbitmq
}

// Create MQ instance with routing mode
func NewRabbitMQRouting(queuename, exchangename, routekey string) *RabbitMQ {
	rabbitmq := NewRabbitMQ(queuename, exchangename, routekey)
	return rabbitmq
}

// Create MQ instance with topic mode
func NewRabbitMQTopic(queuename, exchangename, routekey string) *RabbitMQ {
	rabbitmq := NewRabbitMQ(queuename, exchangename, routekey)
	return rabbitmq
}

// Producer MQ with simple queue mode
func (mq *RabbitMQ) ProduceMQSimple(message []string) {
	for _, msg := range message {
		//println(msg)
		_, err := mq.channelname.QueueDeclare(
			mq.Queuename,
			false, //durable:
			false, //AutoDelete:
			false, //exclusive:
			false, //noWait:
			nil,   //args:
		)
		if err != nil {
			fmt.Println(err)
		}

		mq.channelname.Publish(
			mq.Exchangename,
			mq.Queuename,
			false, //mandatory:
			false, //immediate:
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(msg),
			},
		)
	}

}

// Consume MQ with simple queue mode
func (mq *RabbitMQ) ConsumeMQSimple() {
	_, err := mq.channelname.QueueDeclare(
		mq.Queuename,
		false, //durable:
		false, //AutoDelete:
		false, //exclusive:
		false, //noWait:
		nil,   //args:
	)
	if err != nil {
		fmt.Println(err)
	}
	msgs, err := mq.channelname.Consume(
		mq.Queuename,
		"",    //consumer:
		true,  //autoAck:
		false, //exclusive:
		false, //noLocal:
		false, //noWait:
		nil,   //args:
	)
	if err != nil {
		fmt.Println(err)
	}
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			s := common.BytesToString(&d.Body)
			fmt.Printf("receve msg is :%s\n", *s)
		}
	}()
	log.Printf(" Please wait...")
	<-forever

}

// Producer MQ with Publish mode
func (mq *RabbitMQ) ProduceMQPubSub(message []string) {
	for _, msg := range message {
		println(msg)
		err := mq.channelname.ExchangeDeclare(
			mq.Exchangename,
			"fanout",
			true,
			false,
			false,
			false,
			nil,
		)
		mq.FailOnError(err, "Failed to declare an excha"+
			msg)
		mq.channelname.Publish(
			mq.Exchangename,
			"",
			false, //mandatory:
			false, //immediate:
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(msg),
			},
		)

	}
}

// Consume MQ with Publish mode
func (mq *RabbitMQ) ConsumeMQPubSub() {
	err := mq.channelname.ExchangeDeclare(
		mq.Exchangename,
		"fanout",
		true,  //durable:
		false, //AutoDelete:
		false, //exclusive:
		false, //noWait:
		nil,   //args:
	)
	if err != nil {
		fmt.Println(err)
	}
	q, err := mq.channelname.QueueDeclare(
		mq.Queuename, //consumer:
		true,         //autoAck:
		false,        //exclusive:
		false,        //noLocal:
		false,        //noWait:
		nil,          //args:
	)
	mq.FailOnError(err, "Failed to declare a queue")
	err = mq.channelname.QueueBind(
		q.Name,
		"",
		mq.Exchangename,
		false,
		nil,
	)
	msgs, err := mq.channelname.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			s := common.BytesToString(&d.Body)
			fmt.Printf("receve msg from exchange %s is :%s\n", mq.Exchangename, *s)
		}
	}()
	log.Printf(" Please wait...,Quit with CTRL+C\n")
	<-forever

}

// Producer MQ with Routing mode
func (mq *RabbitMQ) ProduceMQRouting(message string) {
	msg := message
	err := mq.channelname.ExchangeDeclare(
		mq.Exchangename,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	mq.FailOnError(err, "Failed to declare an excha"+
		msg)
	mq.channelname.Publish(
		mq.Exchangename,
		mq.Routekey,
		false, //mandatory:
		false, //immediate:
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
}

// Consume MQ with Routing mode
func (mq *RabbitMQ) ConsumeMQRouting() {
	err := mq.channelname.ExchangeDeclare(
		mq.Exchangename,
		"direct",
		true,  //durable:
		false, //AutoDelete:
		false, //exclusive:
		false, //noWait:
		nil,   //args:
	)
	if err != nil {
		fmt.Println(err)
	}
	q, err := mq.channelname.QueueDeclare(
		mq.Queuename, //consumer:
		false,        //autoAck:
		false,        //exclusive:
		true,         //noLocal:
		false,        //noWait:
		nil,          //args:
	)
	mq.FailOnError(err, "Failed to declare a queue")
	err = mq.channelname.QueueBind(
		q.Name,
		mq.Routekey,
		mq.Exchangename,
		false,
		nil,
	)
	msgs, err := mq.channelname.Consume(
		q.Name,
		mq.Routekey,
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			s := common.BytesToString(&d.Body)
			fmt.Printf("receve msg from exchange %s is :%s\n", mq.Exchangename, *s)
		}
	}()
	log.Printf(" Please wait...,Quit with CTRL+C\n")
	<-forever

}

// Producer MQ with Topic mode
func (mq *RabbitMQ) ProduceMQTopic(message string) {
	//for _, msg := range message {
	//	println(msg)
	msg := message
	err := mq.channelname.ExchangeDeclare(
		mq.Exchangename,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	mq.FailOnError(err, "Failed to declare an excha"+
		msg)
	mq.channelname.Publish(
		mq.Exchangename,
		mq.Routekey,
		false, //mandatory:
		false, //immediate:
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)

	//}
}

// Consume MQ with Topic mode
func (mq *RabbitMQ) ConsumeMQTopic() {
	err := mq.channelname.ExchangeDeclare(
		mq.Exchangename,
		"topic",
		true,  //durable:
		false, //AutoDelete:
		false, //exclusive:
		false, //noWait:
		nil,   //args:
	)
	if err != nil {
		fmt.Println(err)
	}
	q, err := mq.channelname.QueueDeclare(
		mq.Queuename, //consumer:
		false,        //autoAck:
		false,        //exclusive:
		true,         //noLocal:
		false,        //noWait:
		nil,          //args:
	)
	mq.FailOnError(err, "Failed to declare a queue")
	err = mq.channelname.QueueBind(
		q.Name,
		mq.Routekey,
		mq.Exchangename,
		false,
		nil,
	)
	msgs, err := mq.channelname.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			s := common.BytesToString(&d.Body)
			fmt.Printf("receve msg from exchange %s is :%s\n", mq.Exchangename, *s)
		}
	}()
	log.Printf(" Please wait...,Quit with CTRL+C\n")
	<-forever

}
