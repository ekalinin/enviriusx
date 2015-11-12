package ls

import (
	"fmt"

	"github.com/ekalinin/enviriusx/commands"
	"github.com/ekalinin/enviriusx/env"
)

type LsCmd struct{}

func (cmd *LsCmd) Run() error {
	envs, err := env.GetEnvList()
	if err != nil {
		return nil
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

func (cmd *LsCmd) GetHelp() string {
	return "List environments"
}

func (cmd *LsCmd) GetArgs() []commands.CommandArg {
	return []commands.CommandArg{}
}

func init() {
	commands.Add("ls", func() commands.Commander {
		return &LsCmd{}
	})
}
