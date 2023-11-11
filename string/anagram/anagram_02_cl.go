//program to check string is anagram where each char from input1 is present in input2 or vice-versa
//in this program, input1 and input2 are going to be command line flags

package main

import (
	"fmt"
	"os"
)

func main(){

	input1 := os.Args[1]
        input2 := os.Args[2]

	fmt.Println("i1=",input1,"\n i2=",input2)

	//create a input1_map
	var input1Map map[rune]int
	//initialize the input1Map
	input1Map = make(map[rune]int)
	
	for _, v := range input1{
		
		i, ok := input1Map[v]
		if ok {
			i++
			input1Map[v] = i
		}
		input1Map[v] = 1
	}
	
	for _, v := range input2{
		_, ok := input1Map[v]
		
		if !ok {
			fmt.Println("Not-Anagram")
			return
		}
//		input1Map[]
	}
	fmt.Println(input1,"and", input2, "strings are anagram.")
}
