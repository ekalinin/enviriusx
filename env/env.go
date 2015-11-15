package env

import (
	"github.com/ekalinin/enviriusx/langs"
)

type Environment struct {
	name  string
	langs []langs.Lang
}

func NewEnv(name string) *Environment {
	return &Environment{name, []langs.Lang{}}
}

func (env *Environment) AddLang(l *langs.Lang) {
	env.langs = append(env.langs, *l)
}

func (env *Environment) WriteMeta() {

}

func (env *Environment) IsExists() bool {
	return true
}

func (env *Environment) Create() error {
	return nil
}

func (env *Environment) Remove() error {
	return nil
}

func (env *Environment) GetName() string {
	return env.name
}
