package main

import (
	"log/slog"

	"github.com/isaquecsilva/posts-rest-api/infra/controller/routes"
)

func main() {
	routes.InitRoutes()
	err := routes.StartServer("localhost:8080")
	slog.Error(err.Error())
}
