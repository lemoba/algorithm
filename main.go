package main

import LinkedList "github.com/lemoba/algorithm/linkedlist"

func main() {
	list := LinkedList.NewSinglyLinkedList()

	list.InsertAtHead(1)
	list.InsertAtHead(2)
	list.InsertAtHead(3)

	list.InsertAtTail(4)
	list.InsertAtTail(5)
	list.InsertAtTail(6)

	list.Print()
}
