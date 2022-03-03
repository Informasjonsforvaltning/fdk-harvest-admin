package main

import "github.com/Informasjonsforvaltning/fdk-harvest-admin/config"

func main() {
	config.LoggerSetup()
	router := config.SetupRouter()
	router.Run(":8080")
}
