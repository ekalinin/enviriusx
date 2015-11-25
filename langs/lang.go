package langs

// Lang version of the certain programming language
type Lang struct {
	Name    string
	Version string
}

// NewLang constructor
func NewLang(name string) *Lang {
	return &Lang{name, ""}
}

func (lang *Lang) GetVersion() string {
	return lang.Version
}

func (lang *Lang) GetName() string {
	return lang.Name
}

func (lang *Lang) SetVersion(v string) {
	lang.Version = v
}

// GetURL returns url for download
func (lang *Lang) GetURL() {

}

// ShowVersion show available versions for install
func (lang *Lang) ShowVersion() {

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
