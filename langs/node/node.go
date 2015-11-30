package node

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/ekalinin/enviriusx/langs"
)

type NodeLang struct {
	langs.Lang
}

func (l *NodeLang) ShowVersions() {
	resp, err := langs.HttpGet("https://nodejs.org/dist")
	if err != nil {
		errors.New(
			fmt.Sprintf("Error during getting versions list: %v", err))
	}

	findVersions := regexp.MustCompile(`"node-v\d*\.\d*\.\d*\.tar.gz"`)
	fmt.Println(findVersions.FindAllString(resp, -1))
}

func init() {
	l := NodeLang{*langs.NewLang("node")}
	langs.Add(l.Name, func() langs.LangBuilder {
		return &l
	})
}
