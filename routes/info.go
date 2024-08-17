// routes/info.go
package routes

import (
    "github.com/gofiber/fiber/v2"
)

func InfoRoute(c *fiber.Ctx) error {
    info := fiber.Map{
        "app":     "Fiber App",
        "version": "1.0.0",
        "author":  "rehanshaikh25",
    }
    return c.JSON(info)
}
