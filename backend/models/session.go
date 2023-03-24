package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	SessionID string    `gorm:"primaryKey" json:"session_id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
}

// override the default gorm create function to set the ID to a uuid
func (sesh *Session) BeforeCreate(tx *gorm.DB) (err error) {
	sesh.SessionID = uuid.New().String()
	return
}
