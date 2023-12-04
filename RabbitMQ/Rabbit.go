package RabbitMQ

import "github.com/streadway/amqp"

const MQURL = "amqp://imoocuser:imoocuser@127.0.0.1:5672/imooc"

type RabbitMQ struct {
	conn *amqp.Connection
}
