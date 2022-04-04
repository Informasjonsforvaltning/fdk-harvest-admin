package handlers

import (
	"context"
	"strings"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/logging"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/service"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func RabbitHandler(msg amqp.Delivery) {
	if msg.Body == nil {
		logrus.Errorf("Unable to create source from rabbit message with key %s, no message body!", msg.RoutingKey)
	}
	service := service.InitService()
	if strings.Contains(msg.RoutingKey, "NewDataSource") {
		err := service.CreateDataSourceFromRabbitMessage(context.Background(), msg.Body)
		if err != nil {
			logging.LogAndPrintError(err)
		}
	} else {
		errors := service.ConsumeReport(context.Background(), msg.RoutingKey, msg.Body)
		for _, err := range errors {
			logging.LogAndPrintError(err)
		}
	}
}
