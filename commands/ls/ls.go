package ls

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ekalinin/enviriusx/commands"
	"github.com/ekalinin/enviriusx/env"
)

type Cmd struct{}

func (cmd *Cmd) Run() error {
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

func (cmd *Cmd) GetHelp() string {
	return "List environments"
}

func (cmd *Cmd) GetArgs() []commands.CommandArg {
	return []commands.CommandArg{}
}

func init() {
	cmd := Cmd{}
	cmdNames := strings.Split(reflect.TypeOf(cmd).PkgPath(), "/")
	cmdName := cmdNames[len(cmdNames)-1]

	commands.Add(cmdName, func() commands.Commander {
		return &cmd
	})
}
