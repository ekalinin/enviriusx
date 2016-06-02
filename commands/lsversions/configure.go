package ls_versions

import (
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/ekalinin/enviriusx/langs"
)

// Run command
func (cmd *Cmd) Run(c *kingpin.ParseContext) error {
	for l, show := range cmd.Langs {
		if *show {
			langs.Langs[l]().ShowVersions()
		}
	}

	return nil
}
