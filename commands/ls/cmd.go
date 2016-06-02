package ls

import (
	"fmt"

	"github.com/ekalinin/enviriusx/env"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Run shows all available environments
func (cmd *Cmd) Run(c *kingpin.ParseContext) error {
	envs, err := env.GetEnvList()
	if err != nil {
		return err
	}

	if len(envs) > 0 {
		fmt.Println("Available environment(s):")
		for _, e := range envs {
			fmt.Println(e.GetName())
		}
	} else {
		fmt.Println("No environments.")
	}

	return nil
}
