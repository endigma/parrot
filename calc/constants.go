package calc

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"gitcat.ca/endigma/parrot/utils"
)

type constant struct {
	Value interface{} `json:"value"`
	Desc  string      `json:"desc"`
	Unit  string      `json:"unit"`
}

type constants map[string]constant

func (c constants) Get(name string) (interface{}, error) {
	entry, found := c[name]

	if !found {
		errorMessage := "No parameter '" + name + "' found."
		return nil, errors.New(errorMessage)
	}

	return entry.Value, nil
}

var c constants = make(constants)

func init() {
	var jsonfile *os.File
	var err error

	for _, path := range []string{utils.UserHomeDir() + "/.config/parrot/constants.json", "constants.json"} {
		jsonfile, err = os.Open(path)
		if err == nil {
			return
		}
	}

	if err != nil {

		if err != nil {
			return
		}
	}

	defer jsonfile.Close()
	byteValue, _ := ioutil.ReadAll(jsonfile)
	json.Unmarshal(byteValue, &c)
}

func SetConstant(k string, v interface{}) {
	newconst := &constant{
		Value: v,
		Desc:  "User created variable",
	}

	c[k] = *newconst
}

func GetConstants() map[string]constant {
	return c
}
