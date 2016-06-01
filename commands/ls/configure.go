package ls

import (
	"github.com/ekalinin/enviriusx/commands"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Cmd command
type Cmd struct {
	Name     string
	ShowMeta bool
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
