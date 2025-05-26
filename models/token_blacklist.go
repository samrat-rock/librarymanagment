package models

import "time"

type TokenBlacklist struct {
	ID        uint      `gorm:"primaryKey"`
	Token     string    `gorm:"uniqueIndex"`
	CreatedAt time.Time
}
