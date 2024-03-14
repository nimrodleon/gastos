package repositories

import (
	"time"

	"github.com/google/uuid"
	"github.com/nimrodleon/gastos/database"
	"github.com/nimrodleon/gastos/models"
	"gorm.io/gorm"
)

type ExpenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository() *ExpenseRepository {
	db, _ := database.Connect()
	return &ExpenseRepository{db: db}
}

func (r *ExpenseRepository) GetAll(fromDate, toDate time.Time, noteFilter string) ([]models.Expense, error) {
	var expenses []models.Expense
	query := r.db.Where("created_at BETWEEN ? AND ?", fromDate, toDate)
	if noteFilter != "" {
		query = query.Where("note LIKE ?", "%"+noteFilter+"%")
	}
	err := query.Find(&expenses).Error
	return expenses, err
}

func (r *ExpenseRepository) GetByID(id uuid.UUID) (*models.Expense, error) {
	var expense models.Expense
	err := r.db.First(&expense, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &expense, nil
}

func (r *ExpenseRepository) Create(expense *models.Expense) error {
	expense.ID = uuid.New()
	return r.db.Create(expense).Error
}

func (r *ExpenseRepository) Update(expense *models.Expense) error {
	return r.db.Omit("CreatedAt").Save(expense).Error
}

func (r *ExpenseRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Expense{}, "id = ?", id).Error
}
