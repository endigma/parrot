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
	jsonFile, err := os.Open("constants.json")

	utils.CheckErr(err)

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

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
