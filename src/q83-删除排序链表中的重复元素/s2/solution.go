package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := ListNode{4,
		&ListNode{5,
			&ListNode{7,
				&ListNode{7,
					&ListNode{8,
						&ListNode{9,
							&ListNode{9, nil}}}}}}}
	re := deleteDuplicates(&head)
	printList(re)
}
func printList(l *ListNode) {
	for l != nil {
		fmt.Printf("%+v  %p\n", l, l)
		l = l.Next
	}
}
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head
	p1 := dummy
	p2 := head

	for p2.Next != nil {
		if p2.Next.Val != p1.Next.Val {
			p1 = p1.Next
			p2 = p2.Next
		} else {
			for p2.Next != nil && p2.Next.Val == p1.Next.Val {
				p2 = p2.Next
			}
			p1.Next.Next = p2.Next
			p1 = p1.Next
			if p2.Next != nil {
				p2 = p2.Next
			}

		}
	}
	return dummy.Next
}
