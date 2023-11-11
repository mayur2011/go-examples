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
	if len(input1) != len(input2){
		fmt.Println(input1,"and", input2, "Strings are Not-Anagram")
		return
	}

	result := "Anagram"


	fmt.Println("i1=",input1,"\ni2=",input2)

	//create a input1_map
	var inputMap map[rune]int
	//initialize the input1Map
	inputMap = make(map[rune]int)
	
	for _, v := range input1{
		
		i, ok := inputMap[v]
		if ok {
			i++
			inputMap[v] = i
			//fmt.Println(inputMap[v])
		}else {
			inputMap[v] = 1
		}	
	}
	tmpMap := inputMap
//	fmt.Println("tmpMap--",tmpMap)
	for _, v := range input2{
		i, ok := tmpMap[v]
		
		if !ok {
			result = "Not-Anagram"
			//return
		}
		if i==0 {
			result = "Not-Anagram"
			//return
		}else {
			i--
			tmpMap[v] = i
		}
	}
	fmt.Println(input1,"and", input2, "Strings are", result)
}
