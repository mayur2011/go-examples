package main

import "fmt"

func main() {

	x := []int{67, 32, 43, 21, 02, 55}
	size := len(x)
	m := size - 1

	for m >= 0 {
		e := 0
		f := e + 1
		for f < m {
			fmt.Printf("%d %d\n", e, f)
			e++
			f++
		}
		m--
	}
}

/*
0 1
1 2
2 3
3 4
4 5
0 1
1 2
2 3
3 4
0 1
1 2
2 3
0 1
1 2
0 1
*/
