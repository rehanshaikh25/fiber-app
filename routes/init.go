package routes

import (
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
    app.Get("/info", InfoRoute)
    app.Get("/members", MembersRoute(db))
    app.Get("/", WelcomeRoute)
}
