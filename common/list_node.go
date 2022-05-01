package common

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func List2Node(nums []int) *ListNode {
	root := &ListNode{
		Val: nums[0],
	}
	node := root
	for i := 1; i < len(nums); i++ {
		node.Next = &ListNode{
			Val: nums[i],
		}
		node = node.Next
	}
	return root
}

func (n *ListNode) Walk() {
	node := n
	for node != nil {
		fmt.Printf("%v\n", node.Val)
		node = node.Next
	}
}
