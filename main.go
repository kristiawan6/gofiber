package main

import (
    "fetchAPI_gofiber/src/config"
    "fetchAPI_gofiber/src/helper"
    "fetchAPI_gofiber/src/routes"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
    app := fiber.New()

    // Initialize database and perform migrations
    config.InitDB()
    helper.Migration()

    // Add CORS middleware
    app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
        AllowHeaders: "Content-Type",
    }))

    // Register routes
    routes.Router(app)

    // Start the server on port 8080
    app.Listen(":8080")
}
