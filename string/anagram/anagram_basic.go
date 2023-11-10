package main

import "fmt"

func main(){
	input1 := "LISTEN"
	//input2 := "SILENT"

	//create a input1_map
	var input1Map map[rune]int

	//initialize the input1Map
	input1Map = make(map[rune]int)
	

	for _, v:= range input1{
		fmt.Printf("%c\n",v)
		a, ok := input1Map[v]
		if ok {
			fmt.Println(a,ok)
		}else {
			fmt.Println("missing:",v)
		}
		//fmt.Println(input1Map)
	}
	fmt.Println(input1Map)

}
