package main

import "fmt"

func main(){
	num0 := [3]int{1,2,3}
	num1 := [3]int{4,5,6}
	num2 := [3]int{7,8,9}

	sl := [][3]int{num0, num1,num2}
	fmt.Println("slice of integer array is --\n",sl)

}

