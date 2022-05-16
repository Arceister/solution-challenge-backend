package models

import (
	"api-solution/lib"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewMigrations),
	fx.Provide(NewMongoCall),
)

type Migrations struct {
	db lib.Database
}

type MongoCall struct {
	mongo lib.Mongo
}

func NewMigrations(db lib.Database) Migrations {
	return Migrations{
		db: db,
	}
}

func NewMongoCall(mongo lib.Mongo) MongoCall {
	return MongoCall{
		mongo: mongo,
	}
}

func (m Migrations) Migrate() {
	m.db.DB.AutoMigrate(&User{}, &Donate{})
}

func (m MongoCall) Mongo() {
	ctx := m.mongo.Ctx
	m.mongo.Mongo.CreateCollection(*ctx, "test-purpose")
}
