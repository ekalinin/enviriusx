package mk

import (
	"github.com/ekalinin/enviriusx/commands"
	"gopkg.in/alecthomas/kingpin.v2"
)

type MkCmd struct {
	Name    string
	EnvName string
}

func (cmd *MkCmd) Run(c *kingpin.ParseContext) error {
	return nil
}

func (c *MkCmd) Configure(app *kingpin.Application) {
	ls := app.Command(c.Name, "Create environment").Action(c.Run)
	ls.Flag("env-name", "Environment name").Short('n').StringVar(&c.EnvName)
}

func init() {
	cmd := MkCmd{"mk", ""}
	commands.Add(cmd.Name, func() commands.Commander {
		return &cmd
	})
}
