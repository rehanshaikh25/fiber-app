package main

import (
    "github.com/gofiber/fiber/v2"
)

func main() {
   app := fiber.New()
  
   // GET API endpoint
   app.Get("/", func(c *fiber.Ctx) error {
      return c.SendString("Hello, this is a GET API")
   })
  
   // POST API endpoint
   app.Post("/", func(c *fiber.Ctx) error {
      return c.SendString("Hello, this is a POST API")
   })
  
   // Start the server
   app.Listen(":3000")
}
