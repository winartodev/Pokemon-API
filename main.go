package main

import (
	controllers "Pokemon-API/Controllers"
	models "Pokemon-API/Models"
)

func main() {
	db := models.Connect()
	defer db.Close()

	controllers.EndpointsHandler()
	controllers.Run()
}