package models

import "time"

type Donate struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	User        []User    `gorm:"many2many:users_donate;"`
	Judul       string    `gorm:"notNull" json:"judul"`
	Gambar      string    `json:"gambar"`
	Deskripsi   string    `gorm:"notNull" json:"deskripsi"`
	JenisProduk string    `json:"jenisProduk"`
	Kuantitas   uint      `gorm:"notNull" json:"kuantitas"`
	Aktivitas   bool      `gorm:"notNull" json:"aktivitas"` //True:
	Distribusi  bool      `json:"distribusi"`
	CreatedAt   time.Time `json:"createdAt"`
	Latitude    float64   `gorm:"type:decimal(10,8);default:0" json:"lat"`
	Longitude   float64   `gorm:"type:decimal(11,8);default:0" json:"lng"`
}
