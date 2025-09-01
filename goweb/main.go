package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi from fiber")
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		type User struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}

		var user User
		c.BodyParser(&user)

		return c.JSON(fiber.Map{
			"message": "User creado",
			"user":    user,
		})
	})

	app.Listen(":3000")
}
