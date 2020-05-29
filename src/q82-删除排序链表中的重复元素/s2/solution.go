package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := ListNode{4, &ListNode{6, &ListNode{6, &ListNode{7, nil}}}}
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
	dummy := new(ListNode)
	dummy.Next = head

	p1 := dummy
	p2 := head
	flag := 0
	for p1.Next != nil {
		for p2.Next != nil && p2.Next.Val == p1.Next.Val {
			p2 = p2.Next
			flag = 1
		}

		if flag == 1 {
			p1.Next = p2.Next

			p2 = p2.Next
			flag = 0
		} else {
			p1 = p2
			if p2.Next != nil {
				p2 = p2.Next
			}
		}
	}
	return dummy.Next

}
