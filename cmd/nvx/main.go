package main

import (
	"os"

	"github.com/ekalinin/enviriusx/commands"
	_ "github.com/ekalinin/enviriusx/commands/loadall"
	_ "github.com/ekalinin/enviriusx/langs/loadall"
	"gopkg.in/alecthomas/kingpin.v2"
)

const version = "0.1.0"

func main() {

	app := kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version(version).Author("Eugene Kalinin")
	kingpin.CommandLine.Help = "Universal Virtual Environments Manager v." + version
	kingpin.CommandLine.HelpFlag.Short('h')

	for _, mkCmd := range commands.Commands {
		cmd := mkCmd()
		cmd.Configure(app)
	}
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
