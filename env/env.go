package env

import (
	"os"
	"path/filepath"

	"github.com/ekalinin/enviriusx/langs"
)

// Environment base class
type Environment struct {
	name  string
	langs []langs.Lang
}

// NewEnv constructor
func NewEnv(name string) *Environment {
	return &Environment{name, []langs.Lang{}}
}

// AddLang adds new language into environment
func (env *Environment) AddLang(l *langs.Lang) {
	env.langs = append(env.langs, *l)
}

// WriteMeta saves meta data into environment's folder
func (env *Environment) WriteMeta() {

}

// IsExists checks if environment exists
func (env *Environment) IsExists() bool {
	_, err := os.Stat(filepath.Join(GetEnvHome(), env.name))
	if err == nil {
		return true //, nil
	}
	if os.IsNotExist(err) {
		return false //, nil
	}
	return true //, err
}

// Create creates environment
func (env *Environment) Create() error {
	return nil
}

// Remove removes environment
func (env *Environment) Remove() error {
	return nil
}

// GetName returns environment's name
func (env *Environment) GetName() string {
	return env.name
}
