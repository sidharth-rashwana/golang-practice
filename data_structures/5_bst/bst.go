package bst

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func Insert(root *Node, val int) *Node {
	if root == nil {
		return &Node{val: val}
	}
	if val < root.val {
		root.left = Insert(root.left, val)
	} else {
		root.right = Insert(root.right, val)
	}
	return root
}

func Inorder(root *Node) {
	if root == nil {
		return
	}
	Inorder(root.left)
	fmt.Print(root.val, " ")
	Inorder(root.right)
}

func main() {
	var root *Node
	root = Insert(root, 10)
	root = Insert(root, 5)
	root = Insert(root, 20)
	root = Insert(root, 15)

	Inorder(root)
	fmt.Println()
}
