package main

import "fmt"

type LinkedList struct {
	head *Node
	len  int
}

type Node struct {
	data string
	next *Node
}

// insert new node always at the end
func (l *LinkedList) Insert(val string) {
	n := Node{}
	n.data = val
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.len++
			return
		}
		ptr = ptr.next
	}

}

func (l LinkedList) Search() int {
	return -1
}

func (l LinkedList) InsertAt() {

}

func (l LinkedList) DeleteAt() {

}

// print all nodes of linked-list
func (l *LinkedList) Print() {
	if l.len == 0 {
		fmt.Println("No nodes in the list")
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		fmt.Println("Node:", ptr)
		ptr = ptr.next
	}
}

func (n *Node) Next() *Node {
	return n.next
}

func main() {
	l := LinkedList{}
	fmt.Println("-----* INSERT A NODE *-----")
	l.Insert("10")
	l.Insert("11")
	l.Insert("14")

	fmt.Println("-----* PRINT LIST *-----")
	l.Print()
}
