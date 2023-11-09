package main

import (
	"fmt"
	"sort"
)

var p = fmt.Printf

func countTriplets(arr []int, s int) int {
	return 0
}

func main() {
	arr := []int{7, 2, 5, 4, 1, 3}
	size := len(arr)
	p("size: %d\n", len(arr))
	//p("size: %d", cap(arr))
	sort.Ints(arr)
	p("%d", arr)
	c := countTriplets(arr, size)
	p("%v", c)
}
