package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Name      string    `gorm:"size:100"`
	Email     string    `gorm:"uniqueIndex"`
	Password  string    `gorm:"not null;type:varchar(255);" json:"senha_hash"`
	CreatedAt time.Time
}

type UserOutput struct {
	id            uuid.UUID `gorm:"primaryKey"`
	nome          string    `gorm:"size:100"`
	email         string    `gorm:"uniqueIndex"`
	senha         string
	cadastrado_em time.Time
	//Status        string `gorm:"column:status;not null;type:enum('expired', 'pending', 'logged');"`
}
