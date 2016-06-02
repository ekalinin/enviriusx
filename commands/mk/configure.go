package mk

import (
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/ekalinin/enviriusx/commands"
	"github.com/ekalinin/enviriusx/langs"
)

// Cmd command description
type Cmd struct {
	Name string
	// EnvName new environment's name
	EnvName string
	// Force is it need to re-creat an environment if it exists
	Force bool
	// AutoOn turn on new environment
	AutoOn bool
	// Langs list of langs that needs to be installed into environment
	Langs map[string]*string
}

// generateEnvName generates a new EnvName
func (cmd *Cmd) generateEnvName() {
	if cmd.EnvName != "" {
		return
	}
	for l, v := range cmd.Langs {
		if *v == "" {
			continue
		}
		cmd.EnvName += l + "-" + *v + "-"
	}
	if cmd.EnvName != "" {
		cmd.EnvName = cmd.EnvName[:len(cmd.EnvName)-1]
	}
}

// Configure Set configuration for the command line
func (cmd *Cmd) Configure(app *kingpin.Application) {
	cl := app.Command(cmd.Name, "Create environment").Action(cmd.Run)
	cl.Arg("name", "Environment name").StringVar(&cmd.EnvName)
	cl.Flag("force", "Re-create environment if it already exists").Short('f').BoolVar(&cmd.Force)
	//cl.Flag("on", "Activate environment after installation").BoolVar(&cmd.AutoOn)

	for l := range langs.Langs {
		cmd.Langs[l] = new(string)
		cl.Flag(l, "Install certain version of the "+l).StringVar(cmd.Langs[l])
	}
}

func init() {
	cmd := Cmd{"mk", "", false, false, map[string]*string{}}
	commands.Add(cmd.Name, func() commands.Commander {
		return &cmd
	})
}
