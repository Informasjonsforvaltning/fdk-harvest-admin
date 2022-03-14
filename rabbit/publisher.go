package rabbit

import (
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config/connection"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config/env"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type Publisher interface {
	Publish(routingKey string, data []byte) error
}

type PublisherImpl struct{}

func (p PublisherImpl) Publish(routingKey string, data []byte) error {
	rabbitConnection, err := connection.GetRabbitConnection()
	if err != nil {
		logrus.Error("unable to establish rabbit connection")
		return err
	}
	return rabbitConnection.Channel.Publish(
		env.ConstantValues.RabbitExchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         data,
			DeliveryMode: amqp.Persistent,
		})
}
