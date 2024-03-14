package models

import (
	"time"

	"github.com/google/uuid"
)

type Income struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Amount     float64   `json:"amount"`
	CategoryID uuid.UUID `json:"category_id"`
	Category   Category  `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Note       string    `json:"note"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
