package models

import (
	"github.com/google/uuid"
)

type Category struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Name         string    `json:"name"`
	CategoryType string    `json:"category_type"`
	Note         string    `json:"note"`
}
