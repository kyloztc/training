package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	res := dfs(root, target)
	if res {
		return nil
	}
	return root
}

func dfs(node *TreeNode, target int) bool {
	if node.Left == nil && node.Right == nil {
		return node.Val == target
	}
	leftNil := true
	rightNil := true
	if node.Left != nil {
		leftNil = dfs(node.Left, target)
		if leftNil {
			node.Left = nil
		}
	}
	if node.Right != nil {
		rightNil = dfs(node.Right, target)
		if rightNil {
			node.Right = nil
		}
	}
	if leftNil && rightNil {
		return node.Val == target
	}
	return false
}