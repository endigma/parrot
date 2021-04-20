package cli

import (
	"gitcat.ca/endigma/parrot/calc"

	"github.com/abiosoft/ishell"
)

func Handle(c *ishell.Context) {
	answer := calc.Process(c.Args)
	c.Printf("\n = %v\n\n", answer)
	calc.UpdateAns(answer)
}
