package main

import "fmt"

// Node represents a single element of the linked list
type Node struct {
	data int
	next *Node
}

// LinkedList represents the linked list
type LinkedList struct {
	head *Node
}

// Insert at the end
func (l *LinkedList) Append(value int) {
	newNode := &Node{data: value}

	// for 1st node in a linked list
	if l.head == nil {
		l.head = newNode
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
}

// Insert at the beginning
func (l *LinkedList) Prepend(value int) {
	newNode := &Node{
		data: value,
		next: l.head,
	}
	l.head = newNode
}

// Delete a node by value
func (l *LinkedList) Delete(value int) {
	if l.head == nil {
		return
	}

	// Case 1: head needs to be deleted
	if l.head.data == value {
		l.head = l.head.next
		return
	}

	// slow = previous node
	slow := l.head
	fast := l.head.next

	for fast != nil {
		if fast.data == value {
			// unlink fast
			slow.next = fast.next
			return
		}
		slow = fast
		fast = fast.next
	}
}

// Search for a value
func (l *LinkedList) Search(value int) bool {
	curr := l.head
	for curr != nil {
		if curr.data == value {
			return true
		}
		curr = curr.next
	}
	return false
}

// Print the linked list
func (l *LinkedList) Print() {
	curr := l.head
	for curr != nil {
		fmt.Printf("%d -> ", curr.data)
		curr = curr.next
	}
	fmt.Println("nil")
}

// Reverse the linked list
func (l *LinkedList) Reverse() {
	var prev *Node
	curr := l.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	l.head = prev
}


// Main function
func main() {
	list := LinkedList{}

	// Insert elements
	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Prepend(5)

	fmt.Println("Initial list:")
	list.Print()

	// Delete an element
	list.Delete(20)
	fmt.Println("After deleting 20:")
	list.Print()

	// Search elements
	fmt.Println("Search 10:", list.Search(10))
	fmt.Println("Search 40:", list.Search(40))

	// Reverse list
	list.Reverse()
	fmt.Println("After reversing:")
	list.Print()
}
