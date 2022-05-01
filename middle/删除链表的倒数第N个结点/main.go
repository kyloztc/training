package main

import (
	"training/common"
)

func removeNthFromEnd(head *common.ListNode, n int) *common.ListNode {
	pre := new(common.ListNode)
	root := new(common.ListNode)
	pre.Next = root
	root.Next = head
	lastNode := root
	for lastNode != nil && n > 0 {
		lastNode = lastNode.Next
		n -= 1
	}

	//fmt.Printf("lastNode: %v\n", lastNode.Val)
	removeNode := pre
	for lastNode != nil {
		lastNode = lastNode.Next
		removeNode = removeNode.Next
	}
	//fmt.Printf("removeNode: %v\n", removeNode.Val)
	removeNode.Next = removeNode.Next.Next
	return root.Next
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	n := 2
	nodeList := common.List2Node(nums)
	res := removeNthFromEnd(nodeList, n)
	res.Walk()
}
