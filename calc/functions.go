package calc

import (
	"fmt"
	"math"

	"gitcat.ca/endigma/parrot/utils"
	"github.com/Knetic/govaluate"
)

var functions map[string]govaluate.ExpressionFunction = make(map[string]govaluate.ExpressionFunction)

func init() {
	functions["sqrt"] = func_sqrt
	functions["sum"] = func_sum
}

func func_sqrt(args ...interface{}) (interface{}, error) {
	return math.Sqrt(utils.ToNum(args[0])), nil
}

func func_sum(args ...interface{}) (interface{}, error) {
	var sum float64 = 0
	for i := range args {
		sum += utils.ToNum(args[i])
		fmt.Println(i)
	}
	return sum, nil
}
