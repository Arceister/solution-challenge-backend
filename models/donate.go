package models

import "time"

type Donate struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Users       []User `gorm:"many2many:users_donate;"`
	Judul       string `gorm:"notNull"`
	Deskripsi   string `gorm:"notNull"`
	JenisProduk string
	Kuantitas   uint `gorm:"notNull"`
	Aktivitas   bool `gorm:"notNull"`
	CreatedAt   time.Time
}
