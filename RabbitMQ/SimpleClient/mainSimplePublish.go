package main

import (
	"fmt"
	"learn_rabbitmq/RabbitMQ-Main"
)

func main() {
	rabbitmq := RabbitMQ_Main.NewRabbitMQSimple("imoocSimple")
	rabbitmq.PublishSimple("Hello imooc!")
	fmt.Println("发送成功！")
}
