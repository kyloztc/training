package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	ptr1 := list1
	ptr2 := list1
	prePtr1 := new(ListNode)
	prePtr1.Next = ptr1
	for i := a; i < b; i++ {
		ptr2 = ptr2.Next
	}
	for i := 0; i < a; i++ {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
		prePtr1 = prePtr1.Next
	}
	prePtr1.Next = list2
	ptr3 := list2
	for ptr3.Next != nil {
		ptr3 = ptr3.Next
	}
	ptr3.Next = ptr2.Next
	return list1
}

func main() {

}
