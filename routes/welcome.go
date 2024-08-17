package routes

import "github.com/gofiber/fiber/v2"

func WelcomeRoute(c *fiber.Ctx) error {
    return c.SendString("Welcome to the Fiber API!")
}
