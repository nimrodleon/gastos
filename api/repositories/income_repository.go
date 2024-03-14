package repositories

import (
	"time"

	"github.com/google/uuid"
	"github.com/nimrodleon/gastos/database"
	"github.com/nimrodleon/gastos/models"
	"gorm.io/gorm"
)

type IncomeRepository struct {
	db *gorm.DB
}

func NewIncomeRepository() *IncomeRepository {
	db, _ := database.Connect()
	return &IncomeRepository{db: db}
}

func (r *IncomeRepository) GetAll(fromDate, toDate time.Time, noteFilter string) ([]models.Income, error) {
	var incomes []models.Income
	query := r.db.Where("created_at BETWEEN ? AND ?", fromDate, toDate)
	if noteFilter != "" {
		query = query.Where("note LIKE ?", "%"+noteFilter+"%")
	}
	err := query.Find(&incomes).Error
	return incomes, err
}

func (r *IncomeRepository) GetByID(id uuid.UUID) (*models.Income, error) {
	var income models.Income
	err := r.db.First(&income, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &income, nil
}

func (r *IncomeRepository) Create(income *models.Income) error {
	income.ID = uuid.New()
	return r.db.Create(income).Error
}

func (r *IncomeRepository) Update(income *models.Income) error {
	return r.db.Omit("CreatedAt").Save(income).Error
}

func (r *IncomeRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Income{}, "id = ?", id).Error
}
