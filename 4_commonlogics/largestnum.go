package main

import (
	"fmt"
)

func getLargestNum(x []int) (int, int) {
	fmt.Printf("%v \n", x)
	length := len(x)
	fmt.Printf("%d", length)

	largest := x[0]
	index := 0
	y := 0
	for y < length {
		if x[y] > largest {
			index = y
			largest = x[y]
		}
		y++
	}
	return largest, index
}

func main() {
	x := []int{60, 345, 890, 6543, 20, 11, 769, 15489, 12000, 110}
	largest, index := getLargestNum(x)
	fmt.Printf("\nLargest value is %d and at index of list[%d]", largest, index)
}
