package main

import (
	"fmt"
)

func main() {

	age := make(map[string]int)
	age["wula"] = 5
	age["haha"] = 1
	fmt.Println(age)
	fmt.Println("wula", age["wula"])
	delete(age, "haha")
	fmt.Println(age)
	delete(age, "keydonotexist")
	fmt.Println(age)
	m := map[string]int{
		"wula": 5,
		"haha": 1,
	}
	fmt.Println(m)

	for key, val := range m {
		fmt.Printf("%v is %v year old", key, val)
	}

	k := keys(m)
	fmt.Println("The key list of the map", k)
	k1 := keysReflect(m)
	fmt.Println("The key list of the map", k1)
}

func keys(inputMap map[string]int) []string {

	keyList := []string{}
	for key, _ := range inputMap {
		keyList = append(keyList, key)
	}
	return keyList
}
func keysReflect(inputMap map[string]interface{}) []interface{} {

	keyList := []interface{}{}
	for key, _ := range inputMap {
		keyList = append(keyList, key)
	}
	return keyList
}
