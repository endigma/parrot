package calc

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"gitcat.ca/endigma/parrot/utils"
)

var constants map[string]interface{} = make(map[string]interface{})

func init() {
	jsonFile, err := os.Open("constants.json")

	utils.CheckErr(err)

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &constants)
}

func SetConstant(k string, v interface{}) {
	constants[k] = v
}

func GetConstants() map[string]interface{} {
	return constants
}
