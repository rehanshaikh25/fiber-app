// routes/members.go
package routes

import (
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

type Member struct {
    ID   uint   `json:"id" gorm:"primaryKey"`
    Name string `json:"name"`
}

// Fetch members from the database
func MembersRoute(db *gorm.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        var members []Member
        result := db.Find(&members)
        if result.Error != nil {
            return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch members"})
        }
        return c.JSON(members)
    }
}
