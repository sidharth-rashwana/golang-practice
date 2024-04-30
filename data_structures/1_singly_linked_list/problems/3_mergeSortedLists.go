package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Merge two sorted linked lists
// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	dummy := &ListNode{}
// 	curr := dummy

// 	for list1 != nil && list2 != nil {
// 		if list1.Val <= list2.Val {
// 			curr.Next = list1
// 			list1 = list1.Next
// 		} else {
// 			curr.Next = list2
// 			list2 = list2.Next
// 		}
// 		curr = curr.Next
// 	}

// 	// Attach remaining nodes
// 	if list1 != nil {
// 		curr.Next = list1
// 	} else {
// 		curr.Next = list2
// 	}

// 	return dummy.Next
// }

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

// Helper: build linked list from slice
func buildList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for _, v := range nums {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

// Helper: print linked list
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
	list1 := buildList([]int{1, 2, 4})
	list2 := buildList([]int{1, 3, 4})

	fmt.Print("list1: ")
	printList(list1)

	fmt.Print("list2: ")
	printList(list2)

	merged := mergeTwoLists(list1, list2)

	fmt.Print("merged: ")
	printList(merged)
}
