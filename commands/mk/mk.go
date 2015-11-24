package mk

import (
	"errors"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/ekalinin/enviriusx/commands"
	"github.com/ekalinin/enviriusx/env"
	"github.com/ekalinin/enviriusx/langs"
)

// Cmd command
type Cmd struct {
	Name    string
	EnvName string
	Force   bool
	AutoOn  bool
	Langs   map[string]*string
}

func (cmd *Cmd) generateEnvName() {
	// TODO: implement
}

// Run command
func (cmd *Cmd) Run(c *kingpin.ParseContext) error {
	// TODO: add  langs itno Env

	if cmd.EnvName == "" {
		cmd.generateEnvName()
	}

	newEnv := env.NewEnv(cmd.EnvName)

	//for l, v := range cmd.Langs {
	//	if v != "" {
	//		newEnv.AddLang(langs.CreateLang(l, v))
	//	}
	//}

	if newEnv.IsExists() {
		if cmd.Force {
			// TODO: add rm commend here with env name
		} else {
			// TODO: check that app error code != 0
			return errors.New("Environment is already exists.")
		}
	}

	newEnv.Create()

	if cmd.AutoOn {
		// TODO: activate environment
	}

	return nil
}

// Configure Set configuration for the command line
func (cmd *Cmd) Configure(app *kingpin.Application) {
	cl := app.Command(cmd.Name, "Create environment").Action(cmd.Run)
	cl.Arg("name", "Environment name").Required().StringVar(&cmd.EnvName)
	cl.Flag("force", "Re-create environment if it already exists").Short('f').BoolVar(&cmd.Force)
	cl.Flag("on", "Activate environment after installation").BoolVar(&cmd.AutoOn)

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
