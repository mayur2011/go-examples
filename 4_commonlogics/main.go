package main

import (
	"fmt"
	"go-examples/4_commonlogics/common"
)

func main() {
	x := []int{60, 345, 890, 6543, 20, 11, 769, 15489, 12000, 110}
	largest, index := common.GetLargestNum(x)
	fmt.Printf("\nLargest value is %d and at index of list[%d]", largest, index)
	common.GetDigitCount(x)
}
