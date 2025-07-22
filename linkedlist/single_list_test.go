package LinkedList

import "testing"

func TestInsertAtHead(t *testing.T) {
	l := NewSinglyLinkedList()
	l.InsertAtHead(10)
	l.InsertAtHead(20)

	if l.head == nil || l.head.Value != 20 {
		t.Errorf("InsertAtHead failed, expected head value 20, got %v", l.head)
	}

	if l.head.Next == nil || l.head.Next.Value != 10 {
		t.Errorf("InsertAtHead failed, second node should be 10, got %v", l.head.Next)
	}
}

func TestInsertAtTail(t *testing.T) {
	l := NewSinglyLinkedList()
	l.InsertAtTail(30)
	l.InsertAtTail(40)

	if l.head == nil || l.head.Value != 30 {
		t.Errorf("InsertAtTail failed, expected head value 30, got %v", l.head)
	}

	if l.head.Next == nil || l.head.Next.Value != 40 {
		t.Errorf("InsertAtTail failed, second node should be 40, got %v", l.head.Next)
	}
}

func TestInsertAtIndex(t *testing.T) {
	l := NewSinglyLinkedList()
	l.InsertAtTail(1)
	l.InsertAtTail(3)

	success := l.InsertAtIndex(1, 2)

	if !success {
		t.Errorf("InsertAtIndex failed as index 1")
	}

	if l.head.Next == nil || l.head.Next.Value != 2 {
		t.Errorf("Expected value 2 at index 1, got %v", l.head.Next)
	}
}

func TestFind(t *testing.T) {
	l := NewSinglyLinkedList()
	l.InsertAtTail(1)
	l.InsertAtTail(3)

	if !l.Find(3) {
		t.Errorf("Find failed to locate existing value 3")
	}

	if l.Find(2) {
		t.Errorf("Find should not locate non-existent value 2")
	}
}

func TestRemoveAtPosition(t *testing.T) {
	l := NewSinglyLinkedList()

	l.InsertAtTail(1)
	l.InsertAtTail(2)
	l.InsertAtTail(3)

	if !l.RemoveAtPosition(2) {
		t.Errorf("RemoveAtPosition failed as index 2")
	}

	if l.head.Next == nil || l.head.Next.Value != 3 {
		t.Errorf("RemoveFirstAtValue failed, expected 1->3")
	}
}

func TestRemoveFirstAtValue(t *testing.T) {
	l := NewSinglyLinkedList()
	l.InsertAtTail(2)
	l.InsertAtTail(1)
	l.InsertAtTail(3)

	if !l.RemoveFirstAtValue(1) {
		t.Errorf("RemoveFirstAtValue failed")
	}

	if l.head.Next == nil || l.head.Next.Value != 3 {
		t.Errorf("RemoveFirstAtValue failed, expected 2->3")
	}
}
