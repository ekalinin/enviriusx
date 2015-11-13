package ls

import (
	"fmt"

	"github.com/ekalinin/enviriusx/commands"
	"github.com/ekalinin/enviriusx/env"
	"gopkg.in/alecthomas/kingpin.v2"
)

type LsCmd struct {
	Name     string
	ShowMeta bool
}

func (cmd *LsCmd) Run(c *kingpin.ParseContext) error {
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

func (c *LsCmd) Configure(app *kingpin.Application) {
	ls := app.Command(c.Name, "List environments").Action(c.Run)
	ls.Flag("show-meta", "").Short('m').BoolVar(&c.ShowMeta)
}

func init() {
	cmd := LsCmd{"ls", true}
	commands.Add(cmd.Name, func() commands.Commander {
		return &cmd
	})
}
