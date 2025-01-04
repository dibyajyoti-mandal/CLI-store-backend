package handlers

import (
	"github.com/dibyajyoti-mandal/cli-backend/database"
	"github.com/dibyajyoti-mandal/cli-backend/models"
	"github.com/gofiber/fiber/v2"
)

func TestRoot(c *fiber.Ctx) error {
	c.JSON("Root dir .. ")
	return nil
}

func GetAllItems(c *fiber.Ctx) error {
	db := database.DB
	var items []models.Item
	db.Find(&items)
	return c.JSON(items)
}

func GetItem(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var item models.Item
	if err := db.First(&item, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Item not found",
		})
	}
	return c.JSON(item)
}

func NewItem(c *fiber.Ctx) error {
	db := database.DB
	item := new(models.Item)
	if err := c.BodyParser(item); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	db.Create(&item)
	return c.JSON(item)
}

func UpdateItem(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var item models.Item

	if err := db.First(&item, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Item not found",
		})
	}

	if item.Available {
		item.Available = false
		if err := db.Save(&item).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Failed to update item availability",
			})
		}
		return c.JSON(fiber.Map{
			"message": "Item availability updated",
			"item":    item,
		})
	}

	return c.Status(400).JSON(fiber.Map{
		"message": "Item is already unavailable",
	})
}

func DeleteItem(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var item models.Item

	if err := db.First(&item, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Item not found",
		})
	}

	if err := db.Delete(&item).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete item",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Item deleted successfully",
	})
}
