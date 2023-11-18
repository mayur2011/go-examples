package main

import (
	"fmt"
        "os"
	"strconv"
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

func (n *Node) removeNodeFromEnd(num int) {

	//get the length of linked list using l varibale
	//position of node to remove.. from the start

	if num > l {
		fmt.Println("Invalid")
	}else if l == num {
	  start = start.next
	}else {
		pos := (l-num)+1
		prev := new(Node)
		curNode := start
		for i:=1; i<=pos; i++ {
			if i == pos {
				prev.next = curNode.next
			}
			prev = curNode
			curNode = curNode.next
		}
	}
}

func main(){
		if len(os.Args) < 2 {
			fmt.Println("looking for command line arguments /input like go run ... 1 2 etc !")
			os.Exit(0)
		}
		
		nthNode := os.Args[1]
		fmt.Println("nth node=",nthNode)

		n:= &Node{}
		s:=[]int{23,28,10,5,67,39,70}
		for i:=0; i<len(s); i++ {
			a:= s[i]
			fmt.Printf("adding data a= %v\n",a)
			n.add(strconv.Itoa(a))
		}
		n.traverse()
		l = n.getLengthOfList()
		fmt.Println("lenght:",l)
		num, _ := strconv.Atoi(nthNode)
		if num >= 1{
			n.removeNodeFromEnd(num)
		}
		n.traverse()
}
