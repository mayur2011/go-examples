package main

import "fmt"

type Node struct {
	data int
	next *Node
}

var start *Node

func (n *Node) add(value int) error {

	if start == nil {
		n.data=value
		start = n

	}
	fmt.Printf("add..done\nstart is %+v and now next is %+v\n",start,n.next)
	return nil
}

func main(){
	
	if start == nil {
		fmt.Printf("start-->%+v\n",start)
		n := Node{}
		n.add(1)
	}else{
		fmt.Printf("start-->%+v\n","not-nil")
	}

}
