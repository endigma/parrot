package calc

import (
	"strings"

	"github.com/Knetic/govaluate"
)

func Process(args []string) (interface{}, error) {
	expression, err := govaluate.NewEvaluableExpressionWithFunctions(strings.Join(args, " "), functions)
	if err != nil {
		return nil, err
	}
	result, err := expression.Eval(c)
	if err != nil {
		return nil, err
	}

	return result, nil
}
