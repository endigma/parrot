package cli

import (
	"strconv"
	"strings"

	"gitcat.ca/endigma/parrot/calc"

	"gitcat.ca/endigma/ishell"
	"github.com/spf13/viper"
)

func Handle(ctx *ishell.Context) {
	var answer interface{}
	var err error

	for i := range ctx.Args {
		if strings.HasPrefix(ctx.Args[i], "0x") || strings.HasPrefix(ctx.Args[i], "0X") {
			ctx.Args[i] = strings.Replace(ctx.Args[i], "0x", "", -1)
			ctx.Args[i] = strings.Replace(ctx.Args[i], "0X", "", -1)

			parsed, err := strconv.ParseInt(ctx.Args[i], 16, 64)
			if err != nil {
				ctx.Println(color_red.Sprint("Non-decimal formats can only be used with whole numbers."))
				return
			}
			ctx.Args[i] = strconv.FormatInt(parsed, 10)
		}

		if strings.HasPrefix(ctx.Args[i], "0b") || strings.HasPrefix(ctx.Args[i], "0B") {
			ctx.Args[i] = strings.Replace(ctx.Args[i], "0b", "", -1)
			ctx.Args[i] = strings.Replace(ctx.Args[i], "0B", "", -1)

			parsed, err := strconv.ParseInt(ctx.Args[i], 2, 64)
			if err != nil {
				ctx.Println(color_red.Sprint("Non-decimal formats can only be used with whole numbers."))
				return
			}
			ctx.Args[i] = strconv.FormatInt(parsed, 10)
		}

		answer, err = calc.Process(ctx.Args)
	}

	if err != nil {
		ctx.Println(err)
	} else {
		switch viper.Get("format.output") {
		case "hex":
			ctx.Printf("\n = %v\n\n", color_green_bold.Sprintf("0x%v", strconv.FormatInt(int64(answer.(float64)), 16)))
		case "binary":
			ctx.Printf("\n = %v\n\n", color_green_bold.Sprintf("0b%v", strconv.FormatInt(int64(answer.(float64)), 2)))
		case "all":
			ctx.Printf("\n = %v\n", color_green_bold.Sprintf("0x%v", strconv.FormatInt(int64(answer.(float64)), 16)))
			ctx.Printf(" = %v\n", color_green_bold.Sprintf("0b%v", strconv.FormatInt(int64(answer.(float64)), 2)))
			ctx.Printf(" = %v\n\n", color_green_bold.Sprintf("%.6v", answer))
		default:
			ctx.Printf("\n = %v\n\n", color_green_bold.Sprintf("%v", answer))
		}

		calc.SetConstant("ans", answer)
	}

}
