package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func evaluateTree(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return root.Val == 1
	}
	leftVal := evaluateTree(root.Left)
	rightVal := evaluateTree(root.Right)
	if root.Val == 2 {
		return leftVal || rightVal
	}
	return leftVal && rightVal
}