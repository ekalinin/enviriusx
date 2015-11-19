package mk

import (
	"errors"

	"github.com/ekalinin/enviriusx/commands"
	"github.com/ekalinin/enviriusx/env"
	"github.com/ekalinin/enviriusx/langs"

	"gopkg.in/alecthomas/kingpin.v2"
)

type MkCmd struct {
	Name    string
	EnvName string
	Force   bool
	AutoOn  bool
	Langs   map[string]*string
}

func (cmd *MkCmd) generateEnvName() {
	// TODO: implement
}

func (cmd *MkCmd) Run(c *kingpin.ParseContext) error {
	// TODO: add  langs itno Env

	if cmd.EnvName == "" {
		cmd.generateEnvName()
	}

	newEnv := env.NewEnv(cmd.EnvName)

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

func (c *MkCmd) Configure(app *kingpin.Application) {
	cmd := app.Command(c.Name, "Create environment").Action(c.Run)
	cmd.Arg("name", "Environment name").Required().StringVar(&c.EnvName)
	cmd.Flag("force", "Re-create environment if it already exists").Short('f').BoolVar(&c.Force)
	cmd.Flag("on", "Activate environment after installation").BoolVar(&c.AutoOn)

	for l, _ := range langs.Langs {
		c.Langs[l] = new(string)
		cmd.Flag(l, "Install certain version of the "+l).StringVar(c.Langs[l])
	}
}

func init() {
	cmd := MkCmd{"mk", "", false, false, map[string]*string{}}
	commands.Add(cmd.Name, func() commands.Commander {
		return &cmd
	})
}
