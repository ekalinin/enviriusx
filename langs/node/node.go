package node

import (
	"github.com/ekalinin/enviriusx/langs"
)

type NodeLang struct {
	langs.Lang
}

func (l *NodeLang) ListVersions() {

}

func init() {
	l := GoLang{"node"}
	langs.Add(cmd.Name, func() langs.LangCreator {
		return &l
	})
}
