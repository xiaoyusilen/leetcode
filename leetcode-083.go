package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head.Next = deleteDuplicates(head.Next)
	if head.Val == head.Next.Val {
		return head.Next
	} else {
		return head
	}
}

func main () {
	nodeA := deleteDuplicates(&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}})
	for ;;  {
		if nodeA == nil {
			break
		}
		fmt.Println(nodeA.Val)
		nodeA = nodeA.Next
	}
}