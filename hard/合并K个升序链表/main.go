package main

import (
	"fmt"
)

func mergeKLists(lists []*ListNode) *ListNode {
	listLen := len(lists)
	if listLen <= 2 {
		return mergeTwoList(lists)
	}
	leftPart := mergeKLists(lists[0 : listLen/2])
	rightPart := mergeKLists(lists[listLen/2:])
	return mergeTwoList([]*ListNode{leftPart, rightPart})
}

func mergeTwoList(list []*ListNode) *ListNode {
	listLen := len(list)
	if listLen == 0 {
		return nil
	}
	if listLen == 1 {
		return list[0]
	}
	node1 := list[0]
	node2 := list[1]
	ans := new(ListNode)
	head := ans
	for node1 != nil && node2 != nil {
		if node1.Val <= node2.Val {
			ans.Next = node1
			node1 = node1.Next
		} else {
			ans.Next = node2
			node2 = node2.Next
		}
		ans = ans.Next
		ans.Next = nil
	}
	if node1 != nil {
		ans.Next = node1
	}
	if node2 != nil {
		ans.Next = node2
	}
	return head.Next
}

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

func main() {
	list1 := []int{1, 4, 5}
	list2 := []int{1, 3, 4}
	list3 := []int{2, 6}
	node1 := List2Node(list1)
	node2 := List2Node(list2)
	node3 := List2Node(list3)
	res := mergeKLists([]*ListNode{node1, node2, node3})
	res.Walk()
}
