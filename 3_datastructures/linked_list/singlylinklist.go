package main

import "fmt"

type Node struct {
	data int
	next *Node
}

var start *Node

func main(){

	fmt.Printf("start-->%+v\n",start)

}
