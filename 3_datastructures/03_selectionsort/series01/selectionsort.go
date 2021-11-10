package main

import "fmt"

func main() {
	x := []int{75, 34, 57, 98, 23, 21, 02}
	size := len(x)
	//fmt.Println(size)
	e := 0
	for e <= (size - 2) {
		//m := e
		f := 1
		for f <= (size - 1) {
			fmt.Printf("%d - %d\n", e, f)
			f++
		}
		e++
	}
}
