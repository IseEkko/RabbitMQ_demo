package main

import (
	"fmt"
	"learn_rabbitmq/RabbitMQ-Main"
)

func main() {
	rabbitmq := RabbitMQ_Main.NewRabbitMQSimple("imoocSimple")
	for i := 0; i < 100; i++ {
		rabbitmq.PublishSimple("Hello imooc!")
		fmt.Println("发送成功！")
	}
}
