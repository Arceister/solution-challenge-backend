package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"not null;uniqueIndex;size:256" json:"email"`
	NoHP      string    `gorm:"not null;uniqueIndex;size:30" json:"no_hp"`
	Password  string    `gorm:"not null" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	XPPoints  uint      `gorm:"default:0" json:"xp_points"`
}
