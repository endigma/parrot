package calc

import (
	"strings"

	"gitcat.ca/endigma/parrot/utils"

	"github.com/Knetic/govaluate"
)

func Process(args []string) interface{} {
	expression, err := govaluate.NewEvaluableExpression(strings.Join(args, " "))

	utils.CheckErr(err)

	result, err := expression.Evaluate(parameters)

	return result
}
