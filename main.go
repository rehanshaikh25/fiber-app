package main

import (
    "fmt"
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com/spf13/viper" // Import viper
    "gorm.io/driver/postgres"
    "gorm.io/gorm"

    "fiber-app/routes"
)

func initDB() (*gorm.DB, error) {
    dsn := viper.GetString("DATABASE_URL") // Use viper to get DATABASE_URL
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
    viper.SetConfigName(".env") // Set the name of the config file
    err := viper.ReadInConfig() // Read the config file
    if err != nil { // Check for errors
        log.Fatal("Error loading .env file")
    }

    app := fiber.New()

    db, err := initDB()
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    routes.SetupRoutes(app, db)

    port := viper.GetInt("PORT") // Use viper to get PORT
    if port == 0 { // Viper returns int type, so check against 0
        port = 3000 // Default port if not set
    }
    log.Fatal(app.Listen(fmt.Sprintf(":%d", port))) // Use %d for integer formatting
}
