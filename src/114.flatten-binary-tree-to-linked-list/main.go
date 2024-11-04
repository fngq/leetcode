package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	ri := root.Right
	flatten(root.Left)
	root.Right = root.Left
	root.Left = nil

	tail := root.Right
	if tail == nil {
		tail = root
	}
	for tail != nil && tail.Right != nil {
		tail = tail.Right
	}
	tail.Right = ri
	flatten(tail.Right)
}

func main() {}
