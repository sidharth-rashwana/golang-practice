package main

import "fmt"

// Node represents a doubly linked list node
type Node struct {
	data int
	prev *Node
	next *Node
}

// DoublyLinkedList structure
type DoublyLinkedList struct {
	head *Node
	tail *Node
}

// Insert at the end
func (l *DoublyLinkedList) Append(value int) {
	newNode := &Node{data: value}

	// Empty list
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	l.tail.next = newNode
	newNode.prev = l.tail
	l.tail = newNode
}

// Insert at the beginning
func (l *DoublyLinkedList) Prepend(value int) {
	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	newNode.next = l.head
	l.head.prev = newNode
	l.head = newNode
}

// Delete by value
func (l *DoublyLinkedList) Delete(value int) {
	if l.head == nil {
		return
	}

	curr := l.head
	for curr != nil && curr.data != value {
		curr = curr.next
	}

	if curr == nil {
		return
	}

	// If head
	if curr == l.head {
		l.head = curr.next
		if l.head != nil {
			l.head.prev = nil
		}
		return
	}

	// If tail
	if curr == l.tail {
		l.tail = curr.prev
		l.tail.next = nil
		return
	}

	// Middle
	curr.prev.next = curr.next
	curr.next.prev = curr.prev
}

// Print forward
func (l *DoublyLinkedList) PrintForward() {
	curr := l.head
	for curr != nil {
		fmt.Printf("%d <-> ", curr.data)
		curr = curr.next
	}
	fmt.Println("nil")
}

// Print backward
func (l *DoublyLinkedList) PrintBackward() {
	curr := l.tail
	for curr != nil {
		fmt.Printf("%d <-> ", curr.data)
		curr = curr.prev
	}
	fmt.Println("nil")
}

// Search
func (l *DoublyLinkedList) Search(value int) bool {
	curr := l.head
	for curr != nil {
		if curr.data == value {
			return true
		}
		curr = curr.next
	}
	return false
}

// Main
func main() {
	list := DoublyLinkedList{}

	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Prepend(5)

	fmt.Println("Forward:")
	list.PrintForward()

	fmt.Println("Backward:")
	list.PrintBackward()

	list.Delete(20)
	fmt.Println("After deleting 20:")
	list.PrintForward()

	fmt.Println("Search 30:", list.Search(30))
}
