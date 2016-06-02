package ls_versions

import "github.com/ekalinin/enviriusx/langs"

func runCmd(cmd *Cmd) error {
	for l, show := range cmd.Langs {
		if *show {
			langs.Langs[l]().ShowVersions()
		}
	}

	return nil
}
