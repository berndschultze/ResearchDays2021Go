package variables

import (
	"encoding/json"
	"io/ioutil"
)

type Variable struct {
	Name string `json:"name"`
	Type int    `json:"type"`
	Path string `json:"path"`
}

func LoadVariables() ([]Variable, error) {
	data := []Variable{}
	file, err := ioutil.ReadFile("config/var_config.json")
	if err != nil {
		return data, err
	}

	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
