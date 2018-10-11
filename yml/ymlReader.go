package yml

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func isMap(v interface{}) bool {
	switch t := v.(type) {
	case map[interface{}]interface{}:
		return true
	default:
		_ = t
		return false
	}
}

func readDict(keyPrefix string, dict map[interface{}]interface{}, result map[string]string) {
	var rkey string = ""
	for key, value := range dict {
		if keyPrefix == "" {
			rkey = key.(string)
		} else {
			rkey = keyPrefix + "." + key.(string)
		}

		// is value is null or blank, set it as blank
		if value == nil {
			result[rkey] = ""
			continue
		}

		if isMap(value) {
			value := value.(map[interface{}]interface{})
			readDict(rkey, value, result)
		} else {
			value := fmt.Sprintf("%v", value)
			result[rkey] = value
		}
	}
}

func Read(filePath string) *map[string]string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Failed to read cfg #%v \n", err)
	}

	c := make(map[interface{}]interface{})
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		e := fmt.Sprintf("yaml format is incorrected, please fix format in %s", filePath)
		panic(e)
	}

	dict := make(map[string]string)

	readDict("", c, dict)

	return &dict
}
