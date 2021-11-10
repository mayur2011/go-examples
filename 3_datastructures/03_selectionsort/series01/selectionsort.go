package main

import "fmt"

func main() {
	x := []int{75, 34, 57, 98, 23, 21, 02}
	size := len(x)
	e := 0
	for e <= (size - 2) {
		m := e
		f := e + 1
		for f <= (size - 1) {
			if x[f] < x[m] {
				m = e //to capture the index of smallest element
			}
			fmt.Printf("%d - %d\n", e, f)
			f++
		}
		g := x[e]
		x[e] = x[m]
		x[m] = g
		e++
	}
}
