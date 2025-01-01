package main

import (
	"fmt"
	"log"

	"github.com/dibyajyoti-mandal/cli-backend/database"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Root endpoint ")
	})
}

func main() {
	fmt.Print("BOOTING SERVER...\n")
	app := fiber.New()
	setupRoutes(app)

	// Connect to db
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	// Close db connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql-DB from gorm.DB: %v", err)
	}
	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
	}()
	// Start the server
	if err := app.Listen(":8000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}

// run with "air server --port 8000"
