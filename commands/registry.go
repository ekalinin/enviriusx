package commands

import "gopkg.in/alecthomas/kingpin.v2"

// Commander inerface for cli command
type Commander interface {
	Run(c *kingpin.ParseContext) error
	Configure(app *kingpin.Application)
}

// Command is a callback type for command handler
type Command func() Commander

// Commands registry of all available commands
// filled by excution function Add
var Commands = map[string]Command{}

// Add is function to add a new function to a Commands map
func Add(name string, cmd Command) {
	Commands[name] = cmd
}
