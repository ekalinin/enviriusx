package golang

import (
	"github.com/ekalinin/enviriusx/langs"
)

type GoLang struct {
	langs.Lang
}

func (l *GoLang) ListVersions() {

}

func init() {
	l := GoLang{"golang"}
	langs.Add(cmd.Name, func() langs.LangCreator {
		return &l
	})
}
