package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port int

func main() {
	flag.IntVar(&port, "port", 3000, "This is doc for port")
	flag.Parse()
	fmt.Printf("the port is %v\n", port)
	err := ServeStatic()
	if err != nil {
		log.Fatalln(err)
	}
}

func ServeStatic() error {
	host := fmt.Sprintf("localhost:%v", port)
	return http.ListenAndServe(host, http.FileServer(http.Dir(".")))
}
