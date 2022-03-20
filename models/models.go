package models

import (
	"api-solution/lib"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewMigrations),
)

type Migrations struct {
	db lib.Database
}

func NewMigrations(db lib.Database) Migrations {
	return Migrations{
		db: db,
	}
}
func (m Migrations) Migrate() {
	m.db.DB.AutoMigrate(&User{}, &Donate{})
}
