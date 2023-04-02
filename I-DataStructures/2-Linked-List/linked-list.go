package main

import (
	"fmt"
)

type Node struct{
	data interface{}
	next *Node
}

type LinkedList struct{
	head *Node
}

func (list *LinkedList) Insert(d interface{}) {
	newNode := &Node{data: d, next: nil}
	if list.head == nil{
		list.head = newNode
	} else {
		pointer := list.head
		for pointer.next != nil {
			pointer = pointer.next
		}
		pointer.next = newNode
	}
}

func (list *LinkedList) Show(){
	pointer := list.head
	for pointer != nil {
		fmt.Printf("%v -> ", pointer.data)
		pointer = pointer.next
	}
	fmt.Print("NULL\n")
}

func main(){
	fmt.Println("Linked Lists in Go")
	// Init the list
	list := LinkedList{}
	list.Insert(22)
	list.Insert(30)
	list.Insert(57)
	list.Show()
}