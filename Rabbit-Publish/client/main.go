package main

import (
	RabbitMQ_Main "learn_rabbitmq/RabbitMQ-Main"
	"strconv"
	"time"
)

func main() {
	rabbitmq := RabbitMQ_Main.NewRabbitMQPubSub("newProducr")
	for i := 0; i < 100; i++ {
		rabbitmq.PublishPub("订阅模式生产者" + strconv.Itoa(i) + "条")
		time.Sleep(1 * time.Second)
	}
}
