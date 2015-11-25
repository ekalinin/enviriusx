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
	l := GoLang{*langs.NewLang("golang")}
	langs.Add(l.Name, func() langs.LangBuilder {
		return &l
	})
}
