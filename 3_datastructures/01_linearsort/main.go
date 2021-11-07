package main

import "fmt"

func main() {
	x := 0
	y := [5]int{}
	//fmt.Printf("Enter values inputs \n")
	fmt.Printf("Enter any 5 numbers :\n")
	for x < 5 {
		fmt.Println(x)
		fmt.Scanf("%d", &y[x])
		x++
	}
	x = 0
	for x < 5 {
		fmt.Printf("Entered valaues are following- \n %d\n", y[x])
		x++
	}
}
