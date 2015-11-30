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
	v := langs.Version
	app := kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version(v).Author("Eugene Kalinin")
	kingpin.CommandLine.Help = "Universal Virtual Environments Manager v" + v
	kingpin.CommandLine.HelpFlag.Short('h')

	for _, mkCmd := range commands.Commands {
		cmd := mkCmd()
		cmd.Configure(app)
	}
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
