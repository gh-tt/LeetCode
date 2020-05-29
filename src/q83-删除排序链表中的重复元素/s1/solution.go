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
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
