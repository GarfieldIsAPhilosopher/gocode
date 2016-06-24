package main

import "fmt"

func main() {

	jobcount := 100
	workercount := 5
	jobs := make(chan int)
	results := make(chan int)
	go func() {
		for i := 0; i < jobcount; i++ {
			jobs <- i
		}

	}()

	for i := 0; i < workercount; i++ {
		go worker(i, jobs, results)
	}
	for i := 0; i < jobcount; i++ {
		<-results
	}
}
func worker(id int, jobs, results chan int) {
	for job := range jobs {
		fmt.Printf("worker %v is doing job %v\n", id, job)
		// time.Sleep(1 * time.Second)
		results <- job
	}

}
