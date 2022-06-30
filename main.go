package main

import (
	"github.com/CharlesVP15/firstAppGO/functions"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// := declarar tipo =
	app := fiber.New()

	// Middleware
	app.Use(logger.New())

	// Servir archivos estáticos
	app.Static("/", "./public")

	app.Get("/", functions.HandleUser)

	// Ejecutarlo en un puerto
	app.Listen(":3000")
}
