package main

import (
	"fmt"
	RabbitMQ_Main "learn_rabbitmq/RabbitMQ-Main"
	"strconv"
	"time"
)

func main() {
	immocONE := RabbitMQ_Main.NewRabbitMQRouting(
		"exImmoc", "ImoocOne")
	immocTWO := RabbitMQ_Main.NewRabbitMQRouting(
		"exImmoc", "ImoocTWO")
	for i := 0; i <= 10; i++ {
		immocONE.PublishRouting("Hellp imooc one" + strconv.Itoa(i))
		immocTWO.PublishRouting("Hello imooc TWO" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
