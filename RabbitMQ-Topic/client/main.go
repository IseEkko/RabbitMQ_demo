package main

import (
	"fmt"
	RabbitMQ_Main "learn_rabbitmq/RabbitMQ-Main"
	"strconv"
	"time"
)

func main() {
	imoocOne := RabbitMQ_Main.NewRabiitMQTopic(
		"exImoocTopic", "imooc.topic.one")
	imoocTwo := RabbitMQ_Main.NewRabiitMQTopic(
		"exImoocTopic", "imooc.topic.two")
	for i := 0; i <= 10; i++ {
		imoocOne.PublishTopic("hello imooc topic one!" + strconv.Itoa(i))
		imoocTwo.PublishTopic("hello imooc topic two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
