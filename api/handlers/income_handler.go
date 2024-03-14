package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/nimrodleon/gastos/models"
	"github.com/nimrodleon/gastos/repositories"
	"github.com/nimrodleon/gastos/utils"
)

func IncomeHandler(app *fiber.App) {
	repo := repositories.NewIncomeRepository()

	incomeGroup := app.Group("/incomes")

	incomeGroup.Get("/", func(c *fiber.Ctx) error {
		fromDateStr := c.Query("from_date", time.Now().AddDate(0, -1, 0).Format("2006-01-02"))
		toDateStr := c.Query("to_date", time.Now().Format("2006-01-02"))

		fromDate, toDate, err := utils.AdjustDateRange(fromDateStr, toDateStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid date",
			})
		}

		noteFilter := c.Query("note_filter", "")
		incomes, err := repo.GetAll(fromDate, toDate, noteFilter)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(incomes)
	})

	incomeGroup.Get("/:id", func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid ID",
			})
		}
		income, err := repo.GetByID(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(income)
	})

	incomeGroup.Post("/", func(c *fiber.Ctx) error {
		income := new(models.Income)
		if err := c.BodyParser(income); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}
		err := repo.Create(income)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(income)
	})

	incomeGroup.Put("/:id", func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid ID",
			})
		}
		income := new(models.Income)
		if err := c.BodyParser(income); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}
		income.ID = id
		err = repo.Update(income)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(income)
	})

	incomeGroup.Delete("/:id", func(c *fiber.Ctx) error {
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
