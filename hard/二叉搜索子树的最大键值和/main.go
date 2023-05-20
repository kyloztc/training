package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans int

func maxSumBST(root *TreeNode) int {
	ans = 0
	isSearchTree(root)
	return ans
}

func isSearchTree(root *TreeNode) (bool, int, int, int) {
	if root == nil {
		return true, 0, 0, 0
	}
	maxValue := root.Val
	minValue := root.Val
	sum := root.Val
	leftBST := true
	rightBST := true
	if root.Left != nil {
		leftBool, leftMax, leftMin, leftSum := isSearchTree(root.Left)
		if leftBool {
			ans = max(ans, leftSum)
		}
		if !leftBool || leftMax >= root.Val {
			leftBST = false
		} else {
			sum += leftSum
			minValue = min(minValue, leftMin)
		}

	}
	if root.Right != nil {
		rightBool, rightMax, rightMin, rightSum := isSearchTree(root.Right)
		if rightBool {
			ans = max(ans, rightSum)
		}
		if !rightBool || rightMin <= root.Val {
			rightBST = false
		} else {
			maxValue = max(maxValue, rightMax)
			sum += rightSum
		}
	}
	if leftBST && rightBST {
		ans = max(ans, sum)
		return true, maxValue, minValue, sum
	}
	return false, 0, 0, 0
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// func maxSumBST(root *TreeNode) int {
// 	treeQueue := make([]*TreeNode, 0)
// 	treeQueue = append(treeQueue, root)
// 	ans := 0
// 	for len(treeQueue) != 0 {
// 		currentNode := treeQueue[0]
// 		treeQueue = treeQueue[1:]
// 		if currentNode.Left != nil {
// 			treeQueue = append(treeQueue, currentNode.Left)
// 		}
// 		if currentNode.Right != nil {
// 			treeQueue = append(treeQueue, currentNode.Right)
// 		}
// 		isBST, _, _, sum := isSearchTree(currentNode)
// 		if isBST {
// 			ans = max(ans, sum)
// 		}
// 	}
// 	return ans
// }
