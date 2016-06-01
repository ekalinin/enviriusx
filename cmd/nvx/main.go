package main

import (
	"os"

	"github.com/ekalinin/enviriusx/commands"
	_ "github.com/ekalinin/enviriusx/commands/loadall"
	"github.com/ekalinin/enviriusx/langs"
	_ "github.com/ekalinin/enviriusx/langs/loadall"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	app := kingpin.
		UsageTemplate(kingpin.CompactUsageTemplate).
		Version(langs.Version).
		Author("Eugene Kalinin")

	kingpin.CommandLine.Help = "Universal Virtual Environments Manager " +
		"v." + langs.Version
	kingpin.CommandLine.HelpFlag.Short('h')

	for _, mkCmd := range commands.Commands {
		mkCmd().Configure(app)
	}

	kingpin.MustParse(app.Parse(os.Args[1:]))
}
