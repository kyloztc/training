package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	nodeList := make([]*TreeNode, 0)
	nodeList = append(nodeList, root)
	nodeCount := 1
	res := 0
	for len(nodeList) != 0 {
		nextNodeCount := 0
		currentCount := 0
		for nodeCount != 0 {
			nodeCount -= 1
			current := nodeList[0]
			currentCount += current.Val
			nodeList = nodeList[1:]
			if current.Left != nil {
				nodeList = append(nodeList, current.Left)
				nextNodeCount += 1
			}
			if current.Right != nil {
				nodeList = append(nodeList, current.Right)
				nextNodeCount += 1
			}
		}
		res = currentCount
		nodeCount = nextNodeCount
	}
	return res
}

func main() {

}
