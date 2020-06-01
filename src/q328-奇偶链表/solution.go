package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, &ListNode{8, nil}}}}}
	printList(oddEvenList(&head))
}
func printList(l *ListNode) {
	for l != nil {
		fmt.Printf("%+v  %p\n", l, l)
		l = l.Next
	}
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	q := head.Next

	tmp := head.Next

	for p.Next != nil && q.Next != nil {
		p.Next = q.Next
		p = p.Next

		q.Next = p.Next
		q = q.Next
	}

	p.Next = tmp
	return head
}
