package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slowPtr := head
	fastPtr := head
	tmpNode := new(ListNode)
	tmpNode.Next = head
	for fastPtr.Next != nil && fastPtr.Next.Next != nil {
		fastPtr = fastPtr.Next.Next
		slowPtr = slowPtr.Next
		tmpNode = tmpNode.Next
	}
	if fastPtr.Next == nil {
		tmpNode.Next = tmpNode.Next.Next
	} else {
		slowPtr.Next = slowPtr.Next.Next
	}
	return head
}
