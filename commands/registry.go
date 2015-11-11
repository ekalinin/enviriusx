package commands

type Commander interface {
	Run() error
	GetHelp() string
	GetDesc() string
}

type Command func() Commander

var Commands = map[string]Command{}

func Add(name string, cmd Command) {
	Commands[name] = cmd
}
