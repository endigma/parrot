package calc

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"gitcat.ca/endigma/parrot/utils"
)

var parameters map[string]interface{} = make(map[string]interface{})

var ans interface{} = 0

func UpdateAns(a interface{}) {
	parameters["ans"] = a
}

func init() {
	jsonFile, err := os.Open("calc/constants.json")

	utils.CheckErr(err)

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &parameters)
}
