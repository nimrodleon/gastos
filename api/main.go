package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nimrodleon/gastos/database"
	"github.com/nimrodleon/gastos/handlers"
	"github.com/nimrodleon/gastos/models"
)

func main() {
	app := fiber.New()
	handlers.CategoryHandler(app)
	handlers.ExpenseHandler(app)

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// AutoMigrate models
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Income{})
	db.AutoMigrate(&models.Expense{})

	app.Listen(":3000")
}
