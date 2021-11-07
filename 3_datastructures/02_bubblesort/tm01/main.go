package main

import "fmt"

func main() {
	x := [7]int{77, 89, 8, 54, 34, 12, 3}
	m := len(x) - 2
	fmt.Printf("Before Bubble sort : \n%d\n", x)
	for m >= 0 {
		e := 0
		f := 1
		for e <= m {
			if x[f] < x[e] {
				g := x[e]
				x[e] = x[f]
				x[f] = g
			}
			e++
			f++
		}
		fmt.Printf("---- %d ----\n", x)
		m--
	}
	fmt.Printf("After Bubble sort : \n%d\n", x)
}
