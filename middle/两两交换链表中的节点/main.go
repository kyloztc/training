package main

import "training/common"

func swapPairs(head *common.ListNode) *common.ListNode {
	root := new(common.ListNode)
	root.Next = head
	ptr := root
	for ptr.Next != nil && ptr.Next.Next != nil {
		first := ptr.Next
		second := ptr.Next.Next
		first.Next = second.Next
		second.Next = first
		ptr.Next = second
		ptr = ptr.Next.Next
	}
	return root.Next
}

func main() {
	nodes := []int{1}
	list := common.List2Node(nodes)
	ans := swapPairs(list)
	ans.Walk()
}
