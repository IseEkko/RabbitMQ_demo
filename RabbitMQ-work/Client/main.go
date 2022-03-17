package main

import (
	"fmt"
	"learn_rabbitmq/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("imoocSimple")
	for i := 0; i < 100; i++ {
		rabbitmq.PublishSimple("Hello imooc!")
		fmt.Println("发送成功！")
	}
}
