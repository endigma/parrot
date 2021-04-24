package cli

import (
	"fmt"

	"gitcat.ca/endigma/parrot/calc"
	"gitcat.ca/endigma/parrot/utils"
	"gitcat.ca/endigma/shella"

	colors "gopkg.in/go-playground/colors.v1"
)

func init() {
	Shell.AddCmd(&shella.Cmd{
		Name: "listconstants",
		Help: "list available numeric constants",
		Handler: func(ctx *shella.Context) {
			constants := calc.GetConstants()

			for c, v := range constants {
				fmt.Printf("%s: %v: %s\n", utils.Color_blue_bold.Sprintf("%s", c), v.Value, v.Desc)
			}
		},
	})

	Shell.AddCmd(&shella.Cmd{
		Name: "set",
		Help: "set <key> <value>",
		Handler: func(ctx *shella.Context) {
			calc.SetConstant(ctx.Args[0], utils.ToNum(ctx.Args[1]))
		},
	})

	Shell.AddCmd(&shella.Cmd{
		Name: "color",
		Help: "color info tool",
		Handler: func(ctx *shella.Context) {
			if len(ctx.Args) != 2 {
				fmt.Println(utils.Color_red.Sprint("Invalid input, please use the format `utils.Colors <utils.Color>`"))
				return
			}

			color, err := colors.Parse(ctx.Args[1])
			if err != nil {
				fmt.Println(utils.Color_red.Sprint("Invalid input, please use the format `utils.Colors <utils.Color>`"))
				return
			}

			fmt.Printf(" %s: %s\n", utils.Color_blue_bold.Sprint("HEX"), color.ToHEX())
			fmt.Printf(" %s: %s\n", utils.Color_blue_bold.Sprint("RGB"), color.ToRGB())
			fmt.Printf(" %s: %s\n", utils.Color_blue_bold.Sprint("RGBA"), color.ToRGBA())
			fmt.Printf(" %s: %t\n", utils.Color_blue_bold.Sprint("Dark"), color.IsDark())
			fmt.Printf(" %s: https://cya.cx/tools/utils.Colors.html?utils.Color=%s\n", utils.Color_blue_bold.Sprint("Preview"), color.ToRGBA())
		},
	})

	Shell.AddCmd(&shella.Cmd{
		Name: "help",
		Help: "print help text for known commands",
		Handler: func(ctx *shella.Context) {
			fmt.Println(utils.Color_blue_bold.Sprint("Commands:"))
			for _, command := range ctx.Shell.Cmds {
				fmt.Printf(" %s: %s\n", utils.Color_blue_bold.Sprint(command.Name), command.Help)
			}
			fmt.Printf("\n%s: %s\n", utils.Color_green_bold.Sprint("Anything else"), "treated as a mathematics query")
		},
	})

	Shell.AddCmd(&shella.Cmd{
		Name: "exit",
		Help: "exit the program",
		Handler: func(ctx *shella.Context) {
			ctx.Shell.Interrupt()
		},
	})
}

// func cmd_output(ctx *shella.Context) {
// 	switch ctx.MultiChoice([]string{
// 		"Decimal",
// 		"Binary",
// 		"Hex",
// 		"All",
// 	}, "Desired output format") {
// 	case 0:
// 		viper.Set("format.output", "decimal")
// 		fmt.Println("Output format set to decimal")
// 	case 1:
// 		viper.Set("format.output", "binary")
// 		fmt.Println("Output format set to binary")
// 	case 2:
// 		viper.Set("format.output", "hex")
// 		fmt.Println("Output format set to hex")
// 	case 3:
// 		viper.Set("format.output", "all")
// 		fmt.Println("Output format set to all")
// 	}
// 	viper.WriteConfig()
// }
