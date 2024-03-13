package models

import (
	"time"

	"github.com/google/uuid"
)

type Expense struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;"`
	Amount     float64
	CategoryID uuid.UUID
	Category   Category `gorm:"foreignKey:CategoryID"`
	Note       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
