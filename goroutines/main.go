package main

import (
	"fmt"
	"time"
)

func main() {

	go say("go")
	go say("for it")
	var in string
	fmt.Scanln(&in)
}

func say(message string) {
	for i := 0; i < 10; i++ {

		time.Sleep(1 * time.Millisecond)
		fmt.Println(message)
	}

}
