package main

import "fmt"

func main() {
	x := [6]int{56, 5, 45, 60, 2, 0}
	fmt.Printf("Before linear sort : \n%d\n", x)
	el := len(x) - 1
	e := 0
	y := 1
	for y <= el {
		f := e + 1
		for f <= el {
			if x[e] > x[f] {
				g := x[e]
				x[e] = x[f]
				x[f] = g
			}
			//fmt.Println(e, "-", f)
			f++
		}
		e++
		y++
	}
	fmt.Printf("After linear sort : \n%d\n", x)
}
