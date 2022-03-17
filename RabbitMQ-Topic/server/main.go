package main

import RabbitMQ_Main "learn_rabbitmq/RabbitMQ-Main"

func main() {
	immocOne := RabbitMQ_Main.NewRabiitMQTopic(
		"exImoocTopic", "#")
	immocOne.RecieveTopic()
}
