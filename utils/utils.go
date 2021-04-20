package utils

import (
	"strconv"

	log "github.com/sirupsen/logrus"
)

func CheckErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func ToNum(v interface{}) float64 {
	switch v.(type) {
	case int:
		return float64(v.(int))
	case float64:
		return v.(float64)
	default:
		var i interface{}

		i, err := strconv.Atoi(v.(string))
		if err != nil {
			i, err = strconv.ParseFloat(v.(string), 64)
			if err != nil {
				log.Fatal(err)
			}
		}
		return i.(float64)
	}
}
