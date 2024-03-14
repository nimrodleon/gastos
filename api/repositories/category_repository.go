package repositories

import (
	"github.com/google/uuid"
	"github.com/nimrodleon/gastos/database"
	"github.com/nimrodleon/gastos/models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository() *CategoryRepository {
	db, _ := database.Connect()
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAll(categoryType string, name string) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Where("category_type = ? AND name LIKE ?", categoryType, "%"+name+"%").Find(&categories).Error
	return categories, err
}

func (r *CategoryRepository) GetByID(id uuid.UUID) (*models.Category, error) {
	var category models.Category
	err := r.db.First(&category, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) Create(category *models.Category) error {
	category.ID = uuid.New()
	return r.db.Create(category).Error
}

func (r *CategoryRepository) Update(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r *CategoryRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Category{}, "id = ?", id).Error
}
