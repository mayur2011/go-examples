package main

import "fmt"

type Node struct {
	data int
	next *Node
}

var start *Node

func (n *Node) add(value int) error {
	
	node:= &Node{data:value}

	if start == nil {
		start = node

	}else{
		curNode := start
		for (curNode.next != nil){
			curNode = curNode.next
		}
		curNode.next = node
	}
	return nil
}

func (n *Node) traverse() {
	curNode := start
	for curNode != nil{
		fmt.Printf("[%d, %p]\n",curNode.data,curNode.next)
		curNode = curNode.next
	}
}

func main(){
		n:= &Node{}
		n.add(1)
		n.add(2)
		n.add(3)
		n.traverse()
}
