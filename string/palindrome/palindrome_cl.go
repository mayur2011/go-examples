package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) < 2 {
		fmt.Println("Plz pass input string like KAYAK after program name")
		os.Exit(1)
	}
	
	s := os.Args[1]

	fmt.Printf("%v\n",s)

	left:= 0
	right:= len(s)-1

	for left <= right {
		//fmt.Printf("%d %c\n",left,s[right])
		if s[left] != s[right]{
			fmt.Println("Not-Palindrome..!")
			return
		}
		left++
		right--
	}
	fmt.Println("String is Palindrome..!")
}
