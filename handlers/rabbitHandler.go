package handlers

import (
	"context"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/logging"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/service"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func RabbitHandler(msg amqp.Delivery) {
	if msg.Body == nil {
		logrus.Error("Unable to create source from rabbit message, no message body!")
	}
	service := service.InitService()
	err := service.CreateDataSourceFromRabbitMessage(context.Background(), msg.Body)
	if err != nil {
		logging.LogAndPrintError(err)
	}
}
