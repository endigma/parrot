package cli

import (
	"os"

	"gitcat.ca/endigma/ishell"
)

var Shell *ishell.Shell = ishell.New()

func Run() {
	Shell.SetHomeHistoryPath(".parrothistory")
	Shell.SetPrompt(color_cyan_bold.Sprint("→ "))
	Shell.Interrupt(func(ctx *ishell.Context, count int, input string) {
		ctx.Println(color_red_bold.Sprint("Bye!"))
		os.Exit(0)
	})
	Shell.NotFound(Handle)

	// when started with "c" as first argument, assume non-interactive execution
	if len(os.Args) > 1 && os.Args[1] == "c" {
		Shell.Process(os.Args[2:]...)
	} else {
		Shell.Printf("Welcome to %s!\n", color_cyan_bold.Sprint("parrot"))
		Shell.Run()
	}
}
