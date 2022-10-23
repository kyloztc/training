package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	slowPtr := head
	fastPtr := head
	for fastPtr.Next != nil && fastPtr.Next.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
	}
	// fmt.Printf("slow: %v\n", slowPtr.Val)
	halfPtr := slowPtr.Next
	slowPtr.Next = nil
	for halfPtr != nil {
		tmp := halfPtr.Next
		halfPtr.Next = slowPtr.Next
		slowPtr.Next = halfPtr
		halfPtr = tmp
	}
	// head.Walk()
	maxValue := 0
	halfPtr = slowPtr.Next
	for halfPtr != nil {
		value := head.Val + halfPtr.Val
		maxValue = max(maxValue, value)
		halfPtr = halfPtr.Next
		head = head.Next
	}
	return maxValue
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	nums := []int{5, 4, 2, 1}
	node := List2Node(nums)
	res := pairSum(node)
	fmt.Printf("res: %v\n", res)
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
