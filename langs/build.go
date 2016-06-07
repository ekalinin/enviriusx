package langs

// Lang version of the certain programming language
type Lang struct {
	Name    string
	Version string
}

// LangVersioner shows all available versions
type LangVersionPresenter interface {
	ShowVersions()
}

// LangVersioner allow work with versions
type LangVersioner interface {
	GetName() string
	GetVersion() string
	SetVersion(string)
}

type LangDeployer interface {
	//GetUrl()

	//CheckDeps() bool
	//Download() error
	//Unpack() error
	//Build() error
	//Configure() error
	//Install() error
	Deploy() error

	//RmSrc() error
	//CopyBin() error
}

// LangBuilder build certain version of the language
type LangBuilder interface {
	LangVersioner
	LangVersionPresenter
	LangDeployer
}

// NewLang constructor
func NewLang(name string) *Lang {
	return &Lang{name, ""}
}

func (lang *Lang) GetName() string {
	return lang.Name
}

func (lang *Lang) GetVersion() string {
	return lang.Version
}

func (lang *Lang) SetVersion(v string) {
	lang.Version = v
}

// CheckDeps checks all deps for language
func (lang *Lang) CheckDeps() error {
	return nil
}

// Download downloads tarball
func (lang *Lang) Download() error {
	return nil
}

// Unpack tarball
func (lang *Lang) Unpack() error {
	return nil
}

// Build call `build`
func (lang *Lang) Build() error {
	return nil
}

// Configure call `configure`
func (lang *Lang) Configure() error {
	return nil
}

// Install call `make install`
func (lang *Lang) Install() error {
	return nil
}

// RmSrc remove source tarball
func (lang *Lang) RmSrc() error {
	return nil
}

// CopyBin copy compiled lang into directory of the virtual environment
func (lang *Lang) CopyBin() error {
	return nil
}

// Deploy full process of the language deployment into environment
func (lang *Lang) Deploy() error {
	if err := lang.CheckDeps(); err != nil {
		return err
	}
	if err := lang.Download(); err != nil {
		return err
	}
	if err := lang.Unpack(); err != nil {
		return err
	}
	if err := lang.Build(); err != nil {
		return err
	}
	if err := lang.Configure(); err != nil {
		return err
	}
	if err := lang.Install(); err != nil {
		return err
	}

	if err := lang.RmSrc(); err != nil {
		return err
	}
	if err := lang.CopyBin(); err != nil {
		return err
	}

	return nil
}
