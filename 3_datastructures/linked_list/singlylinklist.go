package main

import (
	"fmt"
        "os"
)

type Node struct {
	data string
	next *Node
}

var l int
var start *Node

func (n *Node) add(value string) error {
	
	node := &Node{data:value}

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

func (n *Node) getLengthOfList() int{
	i := 0
	if start == nil {
		return i
	}
	
	curNode := start
	for curNode != nil{
		i++
		curNode= curNode.next
	}
	return i
}

func (n *Node) traverse() {
	curNode := start
	for curNode != nil{
		if curNode.next !=nil {
			fmt.Printf("[%v, %p]-->",curNode.data,curNode.next)
		}else{
			fmt.Printf("[%v, %p]\n", curNode.data, curNode.next)
		}

		curNode = curNode.next
	}
}

func (n *Node) getLastNode() *Node{
	curNode:= start
	for curNode != nil{
		if curNode.next == nil {
			return curNode.next
		}
		curNode=curNode.next
	}
	
	return curNode
}

func main(){
		if len(os.Args) < 2 {
			fmt.Println("looking for command line arguments /input like go run ... 1 2 etc !")
			os.Exit(0)
		}
		
		n := &Node{}

		for i:=1; i<len(os.Args); i++ {
			a:= os.Args[i]
			fmt.Printf("adding data i= %v\n",a)
			n.add(a)
		}
		n.traverse()
		ln := n.getLengthOfList()
		fmt.Println("lenght:",ln)
}
