package cli

import (
	"os"

	"gitcat.ca/endigma/ishell"
)

var shell *ishell.Shell = ishell.New()

func Run() {
	shell.SetHomeHistoryPath(".parrothistory")
	shell.SetPrompt(color_cyan_bold.Sprint("→ "))
	shell.Interrupt(func(ctx *ishell.Context, count int, input string) {
		ctx.Println(color_red_bold.Sprint("Bye!"))
		os.Exit(0)
	})
	shell.NotFound(Handle)

	// when started with "c" as first argument, assume non-interactive execution
	if len(os.Args) > 1 && os.Args[1] == "c" {
		shell.Process(os.Args[2:]...)
	} else {
		shell.Printf("Welcome to %s!\n", color_cyan_bold.Sprint("parrot"))
		shell.Run()
	}
}
