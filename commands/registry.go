package commands

import "gopkg.in/alecthomas/kingpin.v2"

type Commander interface {
	Run(c *kingpin.ParseContext) error
	Configure(app *kingpin.Application)
}

type Command func() Commander

var Commands = map[string]Command{}

func Add(name string, cmd Command) {
	Commands[name] = cmd
}
