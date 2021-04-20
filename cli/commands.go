package cli

import (
	"gitcat.ca/endigma/parrot/calc"
	"gitcat.ca/endigma/parrot/utils"
	"github.com/abiosoft/ishell"
)

func init() {
	shell.AddCmd(&ishell.Cmd{
		Name: "listconstants",
		Help: "list numeric constants available to the program",
		Func: cmd_listconstants,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "set",
		Help: "set <key> <value>",
		Func: cmd_set,
	})
}

func cmd_set(ctx *ishell.Context) {
	calc.SetConstant(ctx.Args[0], utils.ToNum(ctx.Args[1]))
}

func cmd_listconstants(ctx *ishell.Context) {
	constants := calc.GetConstants()

	for c, v := range constants {
		ctx.Printf("%s: %v\n", color_blue_bold.Sprintf("%s", c), v)
	}
}
