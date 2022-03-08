package lib

import "os"

type Env struct {
	ServerPort string
}

func NewEnv() Env {
	return Env{}
}

func (env *Env) LoadEnv() {
	env.ServerPort = os.Getenv("PORT")
}
