package main

import (
    "fmt"
    "log"
    "os"

    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"

    "fiber-app/routes"
)

func initDB() (*gorm.DB, error) {
    dsn := os.Getenv("DATABASE_URL")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    err = db.AutoMigrate(&routes.Member{})
    if err != nil {
        return nil, err
    }

    return db, nil
}


func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    app := fiber.New()

    db, err := initDB()
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    routes.SetupRoutes(app, db)

    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }
    log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
