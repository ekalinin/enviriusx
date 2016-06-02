package mk

import (
	"errors"
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/ekalinin/enviriusx/env"
)

// Run command
func (cmd *Cmd) Run(c *kingpin.ParseContext) error {
	cmd.generateEnvName()
	newEnv := env.NewEnv(cmd.EnvName)

	for l, v := range cmd.Langs {
		if *v == "" {
			continue
		}
		newEnv.AddLang(l, *v)
	}

	fmt.Printf(" * Creating environment: %s\n", cmd.EnvName)

	if newEnv.IsExists() {
		if cmd.Force {
			newEnv.Remove()
		} else {
			return errors.New("Environment is already exists")
		}
	}

	newEnv.Create()
	// TODO: activate environment if --on

	return nil
}
