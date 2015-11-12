package mk

import "github.com/ekalinin/enviriusx/commands"

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
	commands.Add("mk", func() commands.Commander {
		return &Cmd{}
	})
}
