package cli

import (
	"strconv"
	"strings"

	"gitcat.ca/endigma/ishell"
	"gitcat.ca/endigma/parrot/calc"
	"gitcat.ca/endigma/parrot/utils"
	"github.com/spf13/viper"

	colors "gopkg.in/go-playground/colors.v1"
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

	shell.AddCmd(&ishell.Cmd{
		Name: "outputf",
		Help: "set the output format",
		Func: cmd_output,
	})

	colors := &ishell.Cmd{
		Name: "colors",
		Help: "tools for working with web colors",
		Func: cmd_colors,
	}

	shell.AddCmd(colors)

}

func cmd_output(ctx *ishell.Context) {
	switch ctx.MultiChoice([]string{
		"Decimal",
		"Binary",
		"Hex",
		"All",
	}, "Desired output format") {
	case 0:
		viper.Set("format.output", "decimal")
		ctx.Println("Output format set to decimal")
	case 1:
		viper.Set("format.output", "binary")
		ctx.Println("Output format set to binary")
	case 2:
		viper.Set("format.output", "hex")
		ctx.Println("Output format set to hex")
	case 3:
		viper.Set("format.output", "all")
		ctx.Println("Output format set to all")
	}
	viper.WriteConfig()
}

func cmd_hextorgb(ctx *ishell.Context) {

	split := strings.Split(strings.Replace(ctx.Args[0], "#", "", -1), "")
	if len(split) != 6 {
		ctx.Println(color_red.Sprint("Invalid input, please use the format `hextorgb ffffff`"))
		return
	}

	var output []int64 = make([]int64, 3)

	rgb := []string{
		strings.Join(split[0:2], ""),
		strings.Join(split[2:4], ""),
		strings.Join(split[4:6], ""),
	}

	for i := range rgb {
		val, err := strconv.ParseInt(rgb[i], 16, 64)
		if err != nil {
			ctx.Println(color_red.Sprint("Invalid input, please use the format `hextorgb ffffff`"))
		}
		output[i] = val
	}

	ctx.Printf(" %s: rgb(%d, %d, %d)\n", color_blue_bold.Sprint("RGB"), output[0], output[1], output[2])
	ctx.Printf(" %s: https://cya.cx/tools/colors.html?color=rgba(%d,%d,%d)\n", color_blue_bold.Sprint("Preview"), output[0], output[1], output[2])
}

func cmd_colors(ctx *ishell.Context) {
	if len(ctx.Args) != 1 {
		ctx.Println(color_red.Sprint("Invalid input, please use the format `colors <color>`"))
		return
	}

	color, err := colors.Parse(ctx.Args[0])
	if err != nil {
		ctx.Println(color_red.Sprint("Invalid input, please use the format `colors <color>`"))
		return
	}

	ctx.Printf(" %s: %s\n", color_blue_bold.Sprint("HEX"), color.ToHEX())
	ctx.Printf(" %s: %s\n", color_blue_bold.Sprint("RGB"), color.ToRGB())
	ctx.Printf(" %s: https://cya.cx/tools/colors.html?color=%s\n", color_blue_bold.Sprint("Preview"), color.ToRGBA())

	// rgb := []string{
	// 	ctx.Args[0],
	// 	ctx.Args[1],
	// 	ctx.Args[2],
	// }

	// var output []int64 = make([]int64, 3)

	// for i := range rgb {
	// 	val, err := strconv.ParseInt(rgb[i], 10, 64)
	//
	// 	if val > 255 {
	// 		ctx.Println(color_red.Sprint("Invalid input, rgb values cannot exceed 255"))
	// 		return
	// 	}
	// 	if val < 0 {
	// 		ctx.Println(color_red.Sprint("Invalid input, rgb values cannot be below 0"))
	// 		return
	// 	}
	// 	output[i] = val
	// }

}

func cmd_set(ctx *ishell.Context) {
	calc.SetConstant(ctx.Args[0], utils.ToNum(ctx.Args[1]))
}

func cmd_listconstants(ctx *ishell.Context) {
	constants := calc.GetConstants()

	for c, v := range constants {
		ctx.Printf("%s: %v: %s\n", color_blue_bold.Sprintf("%s", c), v.Value, v.Desc)
	}
}
