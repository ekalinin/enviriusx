package mk

import (
	"errors"
	"fmt"

	"github.com/ekalinin/enviriusx/env"
)

// Run command
func runCmd(cmd *Cmd) error {
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
