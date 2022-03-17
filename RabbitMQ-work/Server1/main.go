package main

import (
	"learn_rabbitmq/RabbitMQ-Main"
)

func main() {
	rabbitmq := RabbitMQ_Main.NewRabbitMQSimple("imoocSimple")
	rabbitmq.ConsumeSimple()
}
