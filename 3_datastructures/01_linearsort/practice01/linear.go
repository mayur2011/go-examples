package main

import "fmt"

func main() {
	x := []int{45, 32, 76, 01, 21, 98, 01}
	size := len(x) - 1
	e := 0
	for e < size {
		f := e + 1
		for f <= size {
			fmt.Printf("%d %d\n", e, f)
			f++
		}
		e++
	}
}

/*
01
02
03
04
05
06
12
13
14
15
16
..
..
..
56
*/
