package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, &ListNode{8, nil}}}}}
	re := swapPairs(&head)
	printList(re)
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Println(l)
		l = l.Next
	}
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	first := head.Next

	head.Next = swapPairs(first.Next)

	first.Next = head

	return first

}
