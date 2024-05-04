package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// DebugStruct prints the struct in formatted JSON
func DebugStruct(s interface{}) {
	yaml, err := yaml.Marshal(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(yaml))
}

// SaveStruct saves the struct in formatted JSON to a file
func SaveStruct(s interface{}) {
	json, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}

	// save to file '/home/op/Downloads/idea.json'
	err = os.WriteFile("/home/op/Downloads/test.json", json, 0644)
	if err != nil {
		panic(err)
	}
}
