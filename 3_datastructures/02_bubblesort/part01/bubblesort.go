package main

import "fmt"

func main() {
	x := [10]int{76, 43, 56, 32, 21, 12, 98, 02, 5, 60}
	size := len(x) - 2
	for size >= 0 {
		e := 0
		for e <= (size) {
			f := e + 1
			if x[f] < x[e] {
				g := x[e]
				x[e] = x[f]
				x[f] = g
			}
			fmt.Println(e, "-", f)
			e++
		}
		size--
	}
}
