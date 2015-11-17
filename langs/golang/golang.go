package golang

import (
	"github.com/ekalinin/enviriusx/langs"
)

type GoLang struct {
	langs.Lang
}

func (l *GoLang) ShowVersions() {

}

func init() {
	l := GoLang{langs.Lang{"golang"}}
	langs.Add(l.Name, func() langs.LangBuilder {
		return &l
	})
}
