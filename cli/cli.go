package cli

import (
	"github.com/abiosoft/ishell"
	"github.com/fatih/color"
)

var shell *ishell.Shell = ishell.New()

func Run() {
	// create new shell.
	// by default, new shell includes 'exit', 'help' and 'clear' commands.

	cyan := color.New(color.FgCyan).SprintFunc()

	shell.SetPrompt(cyan(">>> "))

	// shell.AddCmd(&ishell.Cmd{
	// 	Name: "precision",
	// 	Help: "set the precision of results, 0 being integers",
	// 	Func: func(c *ishell.Context) {
	// 		convert, err := strconv.Atoi(c.Args[0])
	// 		utils.CheckErr(err)
	// 		SetPrecision(convert)
	// 	},
	// })

	// display welcome info.
	shell.Println("Welcome to gitcat.ca/endigma/parrotshell")

	shell.NotFound(Handle)

	// run shell
	shell.Run()
}
