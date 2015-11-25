package node

import (
	"github.com/ekalinin/enviriusx/langs"
)

type NodeLang struct {
	langs.Lang
}

func (l *NodeLang) ShowVersions() {

}

func init() {
	l := NodeLang{*langs.NewLang("node")}
	langs.Add(l.Name, func() langs.LangBuilder {
		return &l
	})
}
