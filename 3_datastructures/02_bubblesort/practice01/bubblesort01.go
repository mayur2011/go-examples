package main

import "fmt"

func main() {

	x := []int{67, 32, 43, 21, 02, 55}
	fmt.Printf("Before Sorting:%d\n", x)
	size := len(x)
	m := size - 1

	for m >= 0 {
		e := 0
		f := e + 1
		for f <= m {
			if x[f] < x[e] {
				g := x[e]
				x[e] = x[f]
				x[f] = g
			}
			e++
			f++
		}
		m--
	}
	fmt.Printf("After  Sorting: %d", x)
}
