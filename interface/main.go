package main

import "fmt"

type Cat struct {
	name string
}

func (c Cat) Pet() {
	fmt.Println("purrrrrr")

}
func (c Cat) Name() string {
	return c.name
}
func (d Dog) Name() string {
	return d.name
}

type Dog struct {
	name string
}

func (d Dog) Pet() {
	fmt.Println("wangwang")

}

type Animal interface {
	Pet()
	Name() string
}

func Compliment(a Animal) {
	fmt.Println(a.Name())
	a.Pet()
}
func main() {

	c := Cat{name: "garfield"}
	d := Dog{name: "odi"}
	Compliment(c)
	Compliment(d)
}
