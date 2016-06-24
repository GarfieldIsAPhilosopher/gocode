package main

import "fmt"

const PI float64 = 3.14

type point struct {
	x, y int
}

func NewPoint(x, y int) point {
	return point{x, y}

}
func (r rect) area() int {
	return r.width * r.height
}

type rect struct {
	p      point
	width  int
	height int
}
type circle struct {
	radius int
}

func (c circle) diameter() float64 {
	return 2 * PI *float64( c.radius)
}
func main() {
	p := point{10, 20}
	fmt.Println(p)

	rec := rect{
		p:      NewPoint(10, 20),
		width:  10,
		height: 20,
	}
	fmt.Println(rec)
	fmt.Println(rec.area())

	c := circle{
		radius: 10,
	}
	fmt.Println(c.diameter())
}
