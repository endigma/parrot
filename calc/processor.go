package calc

import (
	"strings"

	"github.com/Knetic/govaluate"
)

var functions map[string]govaluate.ExpressionFunction = map[string]govaluate.ExpressionFunction{
	"sqrt": func_sqrt,
}

func Process(args []string) (interface{}, error) {
	expression, err := govaluate.NewEvaluableExpressionWithFunctions(strings.Join(args, " "), functions)
	if err != nil {
		return nil, err
	}
	result, err := expression.Evaluate(constants)
	if err != nil {
		return nil, err
	}

	return result, nil
}
