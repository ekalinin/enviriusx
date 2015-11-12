package commands

type CommandArg struct {
	Name string
	Help string
}

type Commander interface {
	Run() error
	GetHelp() string
	GetArgs() []CommandArg
}

type Command func() Commander

var Commands = map[string]Command{}

func Add(name string, cmd Command) {
	Commands[name] = cmd
}
