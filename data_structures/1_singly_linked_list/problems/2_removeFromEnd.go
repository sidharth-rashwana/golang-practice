package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Dummy node to handle edge cases (like deleting head)
	dummy := &ListNode{Next: head}

	slow := dummy
	fast := dummy

	// Move fast pointer n steps ahead
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// Move both pointers until fast reaches the end
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// Remove the nth node from end
	slow.Next = slow.Next.Next

	return dummy.Next
}

// Helper function to create a linked list from slice
func createList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, v := range nums {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

// Helper function to print the linked list
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	head := createList([]int{1, 2, 3, 4, 5})
	fmt.Print("Original list: ")
	printList(head)

	head = removeNthFromEnd(head, 2)
	fmt.Print("After removing 2nd from end: ")
	printList(head)
}
