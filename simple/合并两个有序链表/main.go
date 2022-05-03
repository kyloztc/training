package main

import "training/common"

func mergeTwoLists(list1 *common.ListNode, list2 *common.ListNode) *common.ListNode {
	head := new(common.ListNode)
	ans := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			ans.Next = list1
			list1 = list1.Next
		} else {
			ans.Next = list2
			list2 = list2.Next
		}
		ans = ans.Next
		ans.Next = nil
	}
	if list1 != nil {
		ans.Next = list1
	}
	if list2 != nil {
		ans.Next = list2
	}
	return head.Next
}

func main() {
	list1 := []int{}
	list2 := []int{1, 3, 4}
	node1 := common.List2Node(list1)
	node2 := common.List2Node(list2)
	ans := mergeTwoLists(node1, node2)
	ans.Walk()
}
