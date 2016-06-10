package lsversions

import (
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/ekalinin/enviriusx/commands"
	"github.com/ekalinin/enviriusx/langs"
)

// Cmd command
type Cmd struct {
	Name  string
	Langs map[string]*bool
}

// Run command
func (cmd *Cmd) Run(c *kingpin.ParseContext) error {
	return runCmd(cmd)
}

// Configure Set configuration for the command line
func (cmd *Cmd) Configure(app *kingpin.Application) {
	lsv := app.Command(cmd.Name, "Lists available versions for certain language").Action(cmd.Run)
	for l := range langs.Langs {
		cmd.Langs[l] = new(bool)
		lsv.Flag(l, "List versions for the "+l).BoolVar(cmd.Langs[l])
	}
}

func init() {
	cmd := Cmd{"ls-versions", map[string]*bool{}}
	commands.Add(cmd.Name, func() commands.Commander {
		return &cmd
	})
}
