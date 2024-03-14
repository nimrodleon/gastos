package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/nimrodleon/gastos/models"
	"github.com/nimrodleon/gastos/repositories"
	"github.com/nimrodleon/gastos/utils"
)

func ExpenseHandler(app *fiber.App) {
	repo := repositories.NewExpenseRepository()

	expenseGroup := app.Group("/expenses")

	expenseGroup.Get("/", func(c *fiber.Ctx) error {
		fromDateStr := c.Query("from_date", time.Now().AddDate(0, -1, 0).Format("2006-01-02"))
		toDateStr := c.Query("to_date", time.Now().Format("2006-01-02"))

		fromDate, toDate, err := utils.ParseAndAdjustDate(fromDateStr, toDateStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid date",
			})
		}

		noteFilter := c.Query("note_filter", "")
		expenses, err := repo.GetAll(fromDate, toDate, noteFilter)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(expenses)
	})

	expenseGroup.Get("/:id", func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid ID",
			})
		}
		expense, err := repo.GetByID(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(expense)
	})

	expenseGroup.Post("/", func(c *fiber.Ctx) error {
		expense := new(models.Expense)
		if err := c.BodyParser(expense); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}
		err := repo.Create(expense)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(expense)
	})

	expenseGroup.Put("/:id", func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid ID",
			})
		}
		expense := new(models.Expense)
		if err := c.BodyParser(expense); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}
		expense.ID = id
		err = repo.Update(expense)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(expense)
	})

	expenseGroup.Delete("/:id", func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid ID",
			})
		}
		err = repo.Delete(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
