package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	ID        string `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	expiresAt time.Time
}

// override the default gorm create function to set the ID to a uuid
func (sesh *Session) BeforeCreate(tx *gorm.DB) (err error) {
	sesh.ID = uuid.New().String()
	return
}
