package connection

import (
	"fmt"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config/env"
	"github.com/streadway/amqp"
)

type RabbitConnection struct {
	Channel *amqp.Channel
}

func rabbitConnectionString() string {
	host := env.RabbitHost()
	port := env.RabbitPort()
	password := env.RabbitPassword()
	user := env.RabbitUsername()

	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%s", user, password, host, port)

	return connectionString
}

func GetRabbitConnection() (RabbitConnection, error) {
	conn, err := amqp.Dial(rabbitConnectionString())
	if err != nil {
		return RabbitConnection{}, err
	}

	ch, err := conn.Channel()
	return RabbitConnection{
		Channel: ch,
	}, err
}
