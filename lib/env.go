package lib

import "os"

type Env struct {
	ServerPort string
	DbUsername string
	DbPassword string
	DbPort     string
	DbName     string
	DbHost     string
}

func NewEnv() Env {
	return Env{}
}

func (env *Env) LoadEnv() {
	env.ServerPort = os.Getenv("PORT")
	env.DbUsername = os.Getenv("DB_USERNAME")
	env.DbPassword = os.Getenv("DB_PASSWORD")
	env.DbPort = os.Getenv("DB_PORT")
	env.DbName = os.Getenv("DB_NAME")
	env.DbHost = os.Getenv("DB_HOST")
}
