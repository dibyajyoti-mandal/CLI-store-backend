package main

import (
	"fmt"
	"log"

	"github.com/dibyajyoti-mandal/cli-backend/database"
	"github.com/dibyajyoti-mandal/cli-backend/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.TestRoot)
	api := app.Group("/api")

	api.Get("/", handlers.GetAllItems)
	api.Get("/item/:id", handlers.GetItem)
	api.Post("/item", handlers.NewItem)
	api.Delete("/item/:id", handlers.DeleteItem)
	api.Put("/item/sell/:id", handlers.UpdateItem)
}

func main() {
	fmt.Print("BOOTING SERVER...\n")
	app := fiber.New()
	app.Use(cors.New())
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
	if err := app.Listen(":8000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}

// run with "air server --port 8000"
