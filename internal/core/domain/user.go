package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Name      string    `gorm:"size:100"`
	Email     string    `gorm:"uniqueIndex"`
	CreatedAt time.Time
}
