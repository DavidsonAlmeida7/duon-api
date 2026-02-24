package domain

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100"`
	Email     string `gorm:"uniqueIndex"`
	CreatedAt time.Time
}
