package main

import (
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/logging"
)

func main() {
	logging.LoggerSetup()
	router := config.SetupRouter()
	router.Run(":8080")
}
