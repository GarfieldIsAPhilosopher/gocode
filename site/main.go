package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mandyozhao/site/pages"
)

func main() {
	layoutPath := flag.String("layout", "layout.html", "layout file path")
	configPath := flag.String("config", "config.json", "config file path")
	if len(os.Args) < 2 {
		log.Fatalln("not enough arguments")
	}
	flag.Parse()
	filename := os.Args[1]
	fmt.Println(*configPath)
	fmt.Println(*layoutPath)
	p, err := pages.NewPage(filename, *configPath)
	if err != nil {
		log.Fatalln(err)
	}
	err = p.Render(*layoutPath, os.Stdout)
	if err != nil {
		log.Fatalln(err)
	}
}
