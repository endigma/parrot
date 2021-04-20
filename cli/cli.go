package cli

import (
	"github.com/abiosoft/ishell"
)

var shell *ishell.Shell = ishell.New()

func Run() {
	shell.SetPrompt(color_cyan_bold.Sprint("→ "))
	shell.Printf("Welcome to %s!\n", color_cyan.Sprint("parrot"))
	shell.NotFound(Handle)
	shell.Run()
}
