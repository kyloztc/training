package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersHelper(l1, l2, 0)
}

func addTwoNumbersHelper(l1 *ListNode, l2 *ListNode, extra int) *ListNode {
	if l1 == nil && l2 == nil {
		if extra != 0 {
			return &ListNode{
				Val:  1,
			}
		}
		return nil
	}
	newNode := new(ListNode)
	val1 := 0
	val2 := 0
	var l1Next *ListNode
	var l2Next *ListNode
	if l1 != nil {
		val1 = l1.Val
		l1Next = l1.Next
	}
	if l2 != nil {
		val2 = l2.Val
		l2Next = l2.Next
	}
	currentValue := val1 + val2 + extra
	newNode.Val = currentValue % 10
	extra = 0
	if currentValue >= 10 {
		extra = 1
	}
	newNode.Next = addTwoNumbersHelper(l1Next, l2Next, extra)
	return newNode
}

func list2Node(nums []int) *ListNode {
	root := &ListNode{
		Val:  nums[0],
	}
	node := root
	for i := 1; i < len(nums); i ++ {
		node.Next = &ListNode{
			Val:  nums[i],
		}
		node = node.Next
	}
	return root
}

func main() {
	num1 := []int{9,9,9,9,9,9,9}
	node1 := list2Node(num1)
	num2 := []int{9,9,9,9}
	node2 := list2Node(num2)
	res := addTwoNumbers(node1, node2)
	for res != nil {
		fmt.Printf("%v\n", res.Val)
		res = res.Next
	}
}
