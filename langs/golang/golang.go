package golang

import (
	"fmt"
	"os/exec"

	"github.com/ekalinin/enviriusx/langs"
)

// GoLang represents golang
type GoLang struct {
	langs.Lang
}

// ShowVersions shows all available versions of golang
func (l *GoLang) ShowVersions() {
	cmd := `curl -s "https://golang.org/dl/" | ` +
		`egrep -o "\/go[0-9]+\.[0-9]+(\.[0-9]+|rc[0-9]+|beta[1-9]+)*\.linux-386.tar.gz" | ` +
		`egrep -o "[0-9]+\.[0-9]+(\.[0-9]+|rc[0-9]+|beta[1-9]+)*" | sort `
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Printf("Failed to execute: %s", cmd)
		fmt.Printf("Error: %s", cmd)
		return
	}
	//println(cmd)
	println(string(out))
}

func init() {
	l := GoLang{*langs.NewLang("golang")}
	langs.Add(l.Name, func() langs.LangBuilder {
		return &l
	})
}
