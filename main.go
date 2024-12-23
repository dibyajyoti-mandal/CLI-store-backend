package main

import (
	"fmt"

	"github.com/dibyajyoti-mandal/cli-backend/database"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	// app.Get("/api/getItems", GetItems)
	// app.Get("/api/getItem", GetItem)
	// app.Post("/api/addItem", AddItem)
	// app.Delete("/api/delItem", DeleteItem)
	fmt.Println("Routes setup...")
}

func initDB() {

}

func main() {
	fmt.Print("BOOTING BACKEND...\n")
	app := fiber.New()
	setupRoutes(app)
	database.Connect()
	app.Listen(":8000")
}
