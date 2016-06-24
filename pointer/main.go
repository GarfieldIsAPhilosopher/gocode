package main

import "fmt"

func main() {
	num := 2
	double(num)
	fmt.Println(num)
	triple(&num)
	fmt.Println(num)
	x := 10
	y := 20
	swap(&x, &y)
	fmt.Println(x)
	fmt.Println(y)
}

func double(a int) {
	a = a * 2
}

func triple(b *int) {
	*b = *b * 3
}

func swap(num1, num2 *int) {
	tmp := *num1
	*num1 = *num2
	*num2 = tmp

}
