package main

import (
	"learn-middleware-example/app/models"
	"learn-middleware-example/app/routes"
)

func main() {
	models.InitDB()

	routes.SetupNewRoutes()
}
