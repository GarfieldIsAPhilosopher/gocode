package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config map[string]interface{}

func (c Config) Name() string {
	return c["name"].(string)
}

// func (c Config) Tags() []string {
// 	strList := []string{}
// 	switch reflect.TypeOf(c["tags"]).Kind() {
// 	case reflect.Slice:
// 		s := reflect.ValueOf(c["tags"])
// 		for i, str := range s {
// 			strList[i] = s.(string)
// 		}
// 		return strList
// 	}

// 	return c["tags"].([]string)
// }
func LoadConfig(path string) (Config, error) {
	var m map[string]interface{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return m, err
	}
	err = json.Unmarshal(data, &m)
	return m, err

}

func main() {
	config, err := LoadConfig("config.json")
	if err != nil {
		panic(err)

	}

	fmt.Println(config)
	fmt.Println(config.Name())
	// fmt.Println(config.Tags())
}
