package routes

import "github.com/gofiber/fiber/v2"

func InitRoutes() {
	app := fiber.New()

	app.Listen(":3000")
}
