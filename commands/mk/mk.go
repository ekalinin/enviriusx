package mk

import (
	"reflect"
	"strings"

	"github.com/ekalinin/enviriusx/commands"
)

type Cmd struct {
}

func (cmd *Cmd) Run() error {
	return nil
}

func (cmd *Cmd) GetHelp() string {
	return "Create environment"
}

func (cmd *Cmd) GetArgs() []commands.CommandArg {
	return []commands.CommandArg{
		{"name", "Environment name to create"},
	}
}

func init() {
	cmd := Cmd{}
	cmdNames := strings.Split(reflect.TypeOf(cmd).PkgPath(), "/")
	cmdName := cmdNames[len(cmdNames)-1]

	commands.Add(cmdName, func() commands.Commander {
		return &cmd
	})
}
