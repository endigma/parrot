package calc

import (
	"math"

	"gitcat.ca/endigma/parrot/utils"
)

func func_sqrt(args ...interface{}) (interface{}, error) {
	return math.Sqrt(utils.ToNum(args[0])), nil
	// switch len(args) {
	// case 1:
	// }
	// return 0
}
