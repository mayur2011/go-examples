package main

import (
	"fmt"
	"sort"
)

var p = fmt.Printf

func main() {
	arr := []int{7, 2, 5, 4, 1, 3}
	p("size: %d\n", len(arr))
	//p("size: %d", cap(arr))
	sort.Ints(arr)
	p("%d", arr)
}
