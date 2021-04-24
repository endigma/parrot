package cli

import (
	"fmt"
	"os"

	"gitcat.ca/endigma/parrot/utils"
	"gitcat.ca/endigma/shella"
)

var Shell *shella.Shell = shella.New()

func Run() {
	Shell.SetHomeHistoryFile(".parrothistory")
	Shell.SetPrompt(utils.Color_cyan_bold.Sprint("→ "))
	Shell.SetInterruptHandler(func() {
		fmt.Println(utils.Color_red_bold.Sprint("Bye!"))
		os.Exit(0)
	})
	Shell.SetHandler(Handler)

	// when started with "c" as first argument, assume non-interactive execution
	if len(os.Args) > 1 && os.Args[1] == "c" {
		Shell.Process(os.Args[2:]...)
	} else {
		fmt.Printf("Welcome to %s!\n", utils.Color_cyan_bold.Sprint("parrot"))
		Shell.Run()
	}
}
