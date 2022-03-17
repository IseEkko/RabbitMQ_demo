package main

import RabbitMQ_Main "learn_rabbitmq/RabbitMQ-Main"

func main() {
	immocIOne := RabbitMQ_Main.NewRabbitMQRouting(
		"exImmoc", "ImoocTWO")
	immocIOne.RecieveRouting()
}
