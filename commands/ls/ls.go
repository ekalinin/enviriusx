package ls

import (
	"fmt"

	"github.com/ekalinin/enviriusx/commands"
	"github.com/ekalinin/enviriusx/env"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Cmd command
type Cmd struct {
	Name     string
	ShowMeta bool
}

// Run command
func (cmd *Cmd) Run(c *kingpin.ParseContext) error {
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

// Configure Set configuration for the command line
func (cmd *Cmd) Configure(app *kingpin.Application) {
	ls := app.Command(cmd.Name, "List environments").Action(cmd.Run)
	ls.Flag("show-meta", "").Short('m').BoolVar(&cmd.ShowMeta)
}

func init() {
	cmd := Cmd{"ls", true}
	commands.Add(cmd.Name, func() commands.Commander {
		return &cmd
	})
}
