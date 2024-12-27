package main

import (
	"fmt"
	"log"

	"github.com/dibyajyoti-mandal/cli-backend/database"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Api test")
	})
	// app.Get("/api/getItems", GetItems)
	// app.Get("/api/getItem", GetItem)
	// app.Post("/api/addItem", AddItem)
	// app.Delete("/api/delItem", DeleteItem)
	fmt.Println("Routes setup...")
}

func main() {
	fmt.Print("BOOTING SERVER...\n")

	// Initialize Fiber app
	app := fiber.New()

	// Set up routes
	setupRoutes(app)

	// Connect to db
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	// Closing db
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB from gorm.DB: %v", err)
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
