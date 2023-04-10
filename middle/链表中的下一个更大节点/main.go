package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	stack := make([]int, 0)
	ans := make([]int, 0)
	nums := make([]int, 0)
	index := 0
	for head != nil {
		nums = append(nums, head.Val)
		ans = append(ans, 0)
		for len(stack) != 0 && nums[stack[len(stack)-1]] < head.Val {
			topIndex := stack[len(stack)-1]
			ans[topIndex] = head.Val
			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, index)
		index += 1
		head = head.Next
	}
	return ans
}
