package models

import (
	"github.com/google/uuid"
)

type Category struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name         string
	CategoryType string
	Note         string
}
