package cli

import (
	"gitcat.ca/endigma/parrot/calc"

	"github.com/abiosoft/ishell"
)

func Handle(c *ishell.Context) {
	answer, err := calc.Process(c.Args)
	if err != nil {
		c.Println(err)
	}
	c.Printf("\n = %v\n\n", color_green_bold.Sprintf("%v", answer))
	calc.SetConstant("ans", answer)
}
