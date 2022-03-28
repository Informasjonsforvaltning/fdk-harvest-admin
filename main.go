package main

import (
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/handlers"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/logging"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/rabbit"
)

func main() {
	logging.LoggerSetup()

	go rabbit.ConsumerImpl{}.StartConsumer(handlers.RabbitHandler)

	router := config.SetupRouter()
	router.Run(":8080")
}
