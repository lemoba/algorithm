package LinkedList

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

type SinglyLinkedList struct {
	head *ListNode
}

// NewSinglyLinkedList 创建链表
func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{head: nil}
}

// InsertAtHead 头插法
func (l *SinglyLinkedList) InsertAtHead(value int) {
	newNode := &ListNode{Value: value}
	newNode.Next = l.head
	l.head = newNode
}

// InsertAtTail 尾插法
func (l *SinglyLinkedList) InsertAtTail(value int) {
	newNode := &ListNode{Value: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	curr := l.head

	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = newNode
}

// InsertAtIndex 按索引插入
func (l *SinglyLinkedList) InsertAtIndex(index int, value int) bool {
	if index < 0 {
		return false
	}

	if index == 0 {
		l.InsertAtHead(value)
	}

	newNode := &ListNode{Value: value}

	curr := l.head

	for i := 0; i < index-1 && curr != nil; i++ {
		curr = curr.Next
	}

	if curr.Next == nil {
		return false
	}

	newNode.Next = curr.Next
	curr.Next = newNode

	return true
}

// Find 按值查找
func (l *SinglyLinkedList) Find(value int) bool {
	curr := l.head

	for curr != nil {
		if curr.Value == value {
			return true
		}
		curr = curr.Next
	}
	return false
}

// RemoveAtPosition 按位置删除从1开始
func (l *SinglyLinkedList) RemoveAtPosition(pos int) bool {
	if pos < 1 || l.head == nil {
		return false
	}

	if pos == 1 {
		l.head = l.head.Next
	}

	curr := l.head

	for i := 1; i < pos-1 && curr != nil; i++ {
		curr = curr.Next
	}

	if curr.Next == nil || curr.Next.Next == nil {
		return false
	}

	curr.Next = curr.Next.Next

	return true
}

// RemoveFirstAtValue 按值删除第一个匹配节点
func (l *SinglyLinkedList) RemoveFirstAtValue(value int) bool {
	if l.head == nil {
		return false
	}

	if l.head.Value == value {
		l.head = l.head.Next
		return true
	}

	curr := l.head

	for curr.Next != nil && curr.Next.Value != value {
		curr = curr.Next
	}

	if curr.Next == nil {
		return false
	}

	curr.Next = curr.Next.Next

	return true
}

// Print 打印链表
func (l *SinglyLinkedList) Print() {
	for curr := l.head; curr != nil; curr = curr.Next {
		fmt.Printf("%d->", curr.Value)
	}
	fmt.Println("nil")
}
