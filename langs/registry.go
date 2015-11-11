package langs

type Lang struct {
	name    string
	version string
}

type LangBuilder interface {
	GetUrl()
	ShowVersion()

	CheckDeps() bool
	Download() error
	Unpack() error
	Build() error
	Configure() error
	Install() error

	RmSrc() error
	CopyBin() error
}

//func mkLang(name string, version string) *Lang {
//	return &Lang(name, version)
//}

func (lang *Lang) CheckDeps() error {
	return nil
}

func (lang *Lang) Download() error {
	return nil
}

func (lang *Lang) Unpack() error {
	return nil
}

func (lang *Lang) Build() error {
	return nil
}

func (lang *Lang) Configure() error {
	return nil
}

func (lang *Lang) Install() error {
	return nil
}

func (lang *Lang) RmSrc() error {
	return nil
}

func (lang *Lang) CopyBin() error {
	return nil
}

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
