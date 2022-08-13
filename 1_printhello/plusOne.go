package main

import "fmt"

// function to add one digit based on diff scenarios
func plusOne(digits []int) []int {

	// initialize an index (i) (digit of units)
	i := len(digits) - 1

	// while the index is valid and the value at [i] ==
	// 9 set it as 0 and move index to previous value
	for i >= 0 && digits[i] == 9 {
		digits[i] = 0
		i--
	}
	if i < 0 {
		//leveraging golang's simplicity with append internal method for array
		return append([]int{1}, digits...)
	} else {
		digits[i]++
	}
	return digits
}

//Driver code
func main() {
	s := []int{9, 9, 9, 9, 9}
	fmt.Println(plusOne(s))
}

// This code is contributed by Mayur

