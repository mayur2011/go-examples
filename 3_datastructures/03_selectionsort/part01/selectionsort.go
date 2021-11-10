package main

import "fmt"

func main() {
	x := []int{75, 34, 57, 98, 23, 02, 21, 13}
	fmt.Printf("Before sorting -- %d\n", x)
	size := len(x)
	e := 0

	for e <= (size - 2) {
		m := e
		f := e + 1
		for f <= (size - 1) {
			if x[f] < x[m] {
				m = f
			}
			//fmt.Printf("%d - %d\n", e, f)
			f++
		}
		g := x[e]
		x[e] = x[m]
		x[m] = g
		e++
	}
	fmt.Printf("After  sorting -- %d\n", x)
}
