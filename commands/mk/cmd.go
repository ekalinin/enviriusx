package mk

import (
	"errors"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/ekalinin/enviriusx/env"
	"github.com/ekalinin/enviriusx/langs"
)

// generateEnvName generates a new EnvName if current is null
func (cmd *Cmd) generateEnvName() {
	if cmd.EnvName != "" {
		return
	}
	for l, v := range cmd.Langs {
		cmd.EnvName += l + "-" + *v + "-"
	}
	if cmd.EnvName != "" {
		cmd.EnvName = cmd.EnvName[:len(cmd.EnvName)-1]
	}
}

// Run command
func (cmd *Cmd) Run(c *kingpin.ParseContext) error {
	if cmd.EnvName == "" {
		cmd.generateEnvName()
	}

	newEnv := env.NewEnv(cmd.EnvName)

	for l, v := range cmd.Langs {
		if *v != "" {
			newLang := langs.Langs[l]()
			newLang.SetVersion(*v)
			newEnv.AddLang(&newLang)
		}
	}

	if newEnv.IsExists() {
		if cmd.Force {
			// TODO: add rm commend here with env name
		} else {
			return errors.New("Environment is already exists.")
		}
	}

	newEnv.Create()

	if cmd.AutoOn {
		// TODO: activate environment
	}

	return nil
}
