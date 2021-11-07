package main

import "fmt"

func main() {
	x := [10]int{76, 54, 60, 10, 23, 34, 21, 12, 22, 05}
	fmt.Println("Initial List:", x)
	size := len(x)
	e := 0
	for e <= (size - 2) {
		f := e + 1
		for f <= (size - 1) {
			if x[f] < x[e] {
				g := x[e]
				x[e] = x[f]
				x[f] = g
			}
			f++
		}
		e++
	}
	y := 0
	for y <= (size - 1) {
		fmt.Printf("%d\n", x[y])
		y++
	}
}
