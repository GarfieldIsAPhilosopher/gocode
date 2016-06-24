package main

import (
	"fmt"
)

func main() {
	var nums [5]int

	fmt.Println("empty:", nums)
	nums[4] = 100
	fmt.Println(nums)
	ints := []int{1, 2, 3, 4, 5}
	fmt.Println("slice", ints)
	ints = append(ints, 6)
	fmt.Println("append", ints)
	fmt.Println("0-2", ints[:2])
	fmt.Println("2-4", ints[2:4])
	fmt.Println("4-6", ints[4:])
	for i, val := range ints {
		fmt.Println(i, val)

	}
	fmt.Println("sum", sum(ints))

}
func sum(list []int) int {
	sum := 0
	for _, val := range list {
		sum = sum + val
	}

	return sum
}
