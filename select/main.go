package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan bool)

	go func() {

		fmt.Println("Running")
		// time.Sleep(100 * time.Millisecond)
		done <- true

	}()
	go func() {

		fmt.Println("Running")
		time.Sleep(1650 * time.Millisecond)
		done <- true

	}()

		time.Sleep(100 * time.Millisecond)
	waiton(done)
	waiton(done)
}
func waiton(done chan bool) {
	select {
	case <-done:
		fmt.Println("Finished")
	case <-time.After(1500 * time.Millisecond):
		fmt.Println("Timeout")
	default:
		fmt.Println("This is the default case")
	}
}
