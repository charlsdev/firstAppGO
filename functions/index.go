package functions

import (
	"github.com/gofiber/fiber/v2"
)

// Definir estructura de JSON - Struct
type User struct {
	Firstname string
	Lastname  string
}

// Función GET
// fiber.Ctx ==> Contexto de Fiber
func HandleUser(c *fiber.Ctx) error {
	user := User{
		Firstname: "Charls",
		Lastname:  "DEV",
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
