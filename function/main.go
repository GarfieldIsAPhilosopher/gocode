package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	fmt.Println(double(5))
	first, _ := parseName("WULA WALA")
	fmt.Println(parseName("wuli wala"))
	greet := func(name string) {
		fmt.Println(name)
	}

	greet(first)

	num := returnFunc()

	fmt.Println(num("12345"))
	logger := log.New(os.Stdout, "mylog", 1)
	logger.Output(2, "tryit")
}

func returnFunc() func(input string) string {

	return func(input string) string {

		return input
	}
}

func double(n int) int {
	return n * 2

}

func parseName(name string) (first, last string) {
	parsed := strings.Split(name, " ")

	return parsed[0], parsed[1]

}
