package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"uniqueIndex;size:256" json:"email"`
	NoHP      string    `gorm:"uniqueIndex;size:30" json:"nohp"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
	XPPoints  uint      `json:"xp"`
}
