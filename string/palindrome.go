package main

import "fmt"

func main(){
	s:= "KAYAK"
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
