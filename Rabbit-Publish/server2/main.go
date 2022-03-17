package main

import RabbitMQ_Main "learn_rabbitmq/RabbitMQ-Main"

func main() {
	rbbitmq := RabbitMQ_Main.NewRabbitMQPubSub("newProducr")
	rbbitmq.RecieveSub()
}
