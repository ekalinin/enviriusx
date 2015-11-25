package langs

// LangBuilder interface
type LangBuilder interface {
	GetName() string
	GetVersion() string
	SetVersion(string)

	ShowVersions()
	//GetUrl()

	//CheckDeps() bool
	//Download() error
	//Unpack() error
	//Build() error
	//Configure() error
	//Install() error

	//RmSrc() error
	//CopyBin() error
}

type LangCreator func() LangBuilder

var Langs = map[string]LangCreator{}

func Add(name string, lc LangCreator) {
	Langs[name] = lc
}
