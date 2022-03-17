package RabbitMQ_Main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

//这里参数：amqp://用户:密码@rabbitmq地址:端口号/host
const MQURL = "amqp://immocuser:immocuser@127.0.0.1:5672/imooc"

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	//队列的名称
	QueueName string
	//交换机
	Exchange string
	//key
	Key string
	//连接信息
	Mqurl string
}

//创建结构体实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     MQURL,
	}
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnerr(err, "创建连接错误！")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnerr(err, "获取channel失败")
	return rabbitmq
}

//断开channel和connection
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

//错误处理函数
func (r *RabbitMQ) failOnerr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

/**
创建简单模式的rabbitmq实例
*/
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName, "", "")
}

//简单模式下生产代码
func (r *RabbitMQ) PublishSimple(message string) {
	//1. 申请队列，如果队列不存在会自动创建，如果存在则跳过创建过程
	//保证队列存在，消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		//是非持久化
		false,
		//是否为自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞
		false,
		//额外的属性
		nil)

	if err != nil {
		fmt.Println(err)
	}
	//2. 发送消息到队列中
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		//如果为true会根据exchange 类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返回给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有绑定消费者，则会把消息还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

func (r *RabbitMQ) ConsumeSimple() {
	//1. 申请队列，如果队列不存在会自动创建，如果存在则跳过创建过程
	//保证队列存在，消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		//是非持久化
		false,
		//是否为自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞
		false,
		//额外的属性
		nil)

	if err != nil {
		fmt.Println(err)
	}
	//接受消息
	msg, err := r.channel.Consume(
		r.QueueName,
		//用于区分多个消费者
		"",
		//是否自动应答
		true,
		//是否具有排他性
		false,
		//如果设置为true，表示不能将同一个connection中发送的消息传递给这个connection的消费者
		false,
		//队列消费是否阻塞
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	forever := make(chan bool)
	//启用协程处理消息
	go func() {
		for d := range msg {
			//实现我们要处理的逻辑函数
			fmt.Printf("Received a message:%s", d.Body)
			fmt.Println(d.Body)
		}
	}()
	log.Println("[*]watting for message ,To exit press CTRL + C")
	<-forever
}

//订阅模式创建RabbitMQ实例
func NewRabbitMQPubSub(exchangeName string) *RabbitMQ {
	//创建RabbitMQ实例
	rabbitmq := NewRabbitMQ("", exchangeName, "")
	var err error
	//获取conntion连接
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnerr(err, "failed to connect rabbitmq!")
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnerr(err, "failed to open a channel")
	return rabbitmq
}

//订阅模式生产
func (r *RabbitMQ) PublishPub(message string) {
	//尝试创建交换机
	err := r.channel.ExchangeDeclare(
		//
		r.Exchange,
		//交换机的类型，这里是广播的类型
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnerr(err, "Failed to declare an excahnge!")
	err = r.channel.Publish(
		r.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

//订阅模式消费端代码
func (r *RabbitMQ) RecieveSub() {
	//试探性创建交换机
	err := r.channel.ExchangeDeclare(
		//
		r.Exchange,
		//交换机的类型，这里是广播的类型
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnerr(err, "Failed to declare an excahnge!")
	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		//排他性
		true,
		false,
		nil)
	r.failOnerr(err, "Failed to declare a queue")

	//绑定队列到exchange中
	err = r.channel.QueueBind(
		//因为这个名字是随机生产的所以这样来取名称/
		q.Name,
		"",
		r.Exchange,
		false,
		nil,
	)

	//消息消费
	message, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	forever := make(chan bool)
	go func() {
		for d := range message {
			log.Printf("Received a message:%s", d.Body)
		}
	}()
	fmt.Println("推出请按 CTRL + C")
	<-forever
}

/**
路由模式
*/
/**
路由模式创建实例
*/
func NewRabbitMQRouting(exchange string, routingKey string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("", exchange, routingKey)
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnerr(err, "Faile to connect rabbitmq!")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnerr(err, "failed to open a channel")
	return rabbitmq
}

//路由模式发送消息
func (r *RabbitMQ) PublishRouting(message string) {
	//尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//注意这里需要改的名称
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnerr(err, "Failed to declare an exchange")
	//2.发送消息
	err = r.channel.Publish(
		r.Exchange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

//路由模式接受消息
func (r *RabbitMQ) RecieveRouting() {
	//1.试探性的创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil)
	r.failOnerr(err, "Failed to declare an excahnge")
	//2.试探性创建队列，这里注意队列的名称不要写
	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil)
	err = r.channel.QueueBind(
		q.Name,
		r.Key,
		r.Exchange,
		false,
		nil)
	message, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	forever := make(chan bool)
	go func() {
		for d := range message {
			log.Printf("Received a message:%s", d.Body)
		}
	}()
	fmt.Println("退出请按CTRL+C")
	<-forever
}
