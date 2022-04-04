package rabbit

import (
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config/connection"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config/env"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type Consumer interface {
	StartConsumer(handler func(d amqp.Delivery)) error
}

type ConsumerImpl struct{}

func (p ConsumerImpl) StartConsumer(handler func(d amqp.Delivery)) {
	rabbitConnection, err := connection.GetRabbitConnection()
	if err != nil {
		logrus.Error(err)
		logrus.Error("unable to establish rabbit connection")
	}

	err = rabbitConnection.Channel.ExchangeDeclare(
		env.ConstantValues.RabbitExchange,
		env.ConstantValues.RabbitExchangeKind,
		false, false, false, false, nil,
	)
	if err != nil {
		logrus.Error(err)
		logrus.Error("unable to declare rabbit exchange")
	}

	_, err = rabbitConnection.Channel.QueueDeclare(
		env.ConstantValues.RabbitListenQueue,
		false, true, false, false, nil,
	)
	if err != nil {
		logrus.Error(err)
		logrus.Error("unable to declare rabbit queue")
	}

	err = rabbitConnection.Channel.QueueBind(
		env.ConstantValues.RabbitListenQueue,
		env.ConstantValues.RabbitNewDataSourceKey,
		env.ConstantValues.RabbitExchange,
		false,
		nil,
	)
	if err != nil {
		logrus.Error(err)
		logrus.Error("unable to bind new data source queue to exchange")
	}

	err = rabbitConnection.Channel.QueueBind(
		env.ConstantValues.RabbitListenQueue,
		env.ConstantValues.RabbitReasonedKey,
		env.ConstantValues.RabbitExchange,
		false,
		nil,
	)
	if err != nil {
		logrus.Error(err)
		logrus.Error("unable to bind reasoned queue to exchange")
	}

	err = rabbitConnection.Channel.QueueBind(
		env.ConstantValues.RabbitListenQueue,
		env.ConstantValues.RabbitHarvestedKey,
		env.ConstantValues.RabbitExchange,
		false,
		nil,
	)
	if err != nil {
		logrus.Error(err)
		logrus.Error("unable to bind harvested queue to exchange")
	}

	err = rabbitConnection.Channel.QueueBind(
		env.ConstantValues.RabbitListenQueue,
		env.ConstantValues.RabbitIngestedKey,
		env.ConstantValues.RabbitExchange,
		false,
		nil,
	)
	if err != nil {
		logrus.Error(err)
		logrus.Error("unable to bind ingested queue to exchange")
	}

	msgs, err := rabbitConnection.Channel.Consume(
		env.ConstantValues.RabbitListenQueue,
		"", false, false, false, false, nil,
	)
	if err != nil {
		logrus.Error(err)
	}

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			handler(msg)
			msg.Ack(false)
		}
	}()
	logrus.Infof("Started listening for messages from exchange %s with queue %s", env.ConstantValues.RabbitExchange, env.ConstantValues.RabbitListenQueue)
	<-forever
}
