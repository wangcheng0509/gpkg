package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Mq struct {
	Host     string
	Port     string
	Username string
	Pwd      string
	Vh       string
	// Queue    string
}

var mqSetting Mq
var ch *amqp.Channel
var conn *amqp.Connection

// var q amqp.Queue

func Init(_mqSetting Mq) {
	mqSetting = _mqSetting
}

func openConn() {
	var err error
	RabbitUrl := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", mqSetting.Username, mqSetting.Pwd, mqSetting.Host, mqSetting.Port, mqSetting.Vh)
	conn, err = amqp.Dial(RabbitUrl)
	failOnError(err, "Failed to connect to RabbitMQ")
	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
}

func SendMsg(exchange, queue string, msg []byte) {
	if conn == nil || conn.IsClosed() || ch == nil {
		openConn()
	}

	err := ch.Publish(
		exchange, //exchange
		queue,    //routing key
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, //Msg set as persistent
			ContentType:  "text/plain",
			Body:         msg,
		})
	failOnError(err, "Failed to publish a message")

}

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s\n", msg, err)
		panic(err)
	}
}
