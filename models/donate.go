package models

import "time"

type Donate struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Users       []User    `gorm:"many2many:users_donate;"`
	Judul       string    `gorm:"notNull" json:"judul"`
	Deskripsi   string    `gorm:"notNull" json:"deskripsi"`
	JenisProduk string    `json:"jenisProduk"`
	Kuantitas   uint      `gorm:"notNull" json:"kuantitas"`
	Aktivitas   bool      `gorm:"notNull" json:"aktivitas"`
	CreatedAt   time.Time `json:"createdAt"`
}
