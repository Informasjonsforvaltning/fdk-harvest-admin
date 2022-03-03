package main

import "github.com/Informasjonsforvaltning/fdk-harvest-admin/config"

func main() {
	router := config.SetupRouter()
	router.Run(":8080")
}
