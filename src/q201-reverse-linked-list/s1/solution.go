package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}
	printList(reverseList(&head))
}
func printList(l *ListNode) {
	for l != nil {
		fmt.Printf("%+v  %p\n", l, l)
		l = l.Next
	}
}

func reverseList(head *ListNode) *ListNode {

	//dummy := new(ListNode)

	//dummy.Next = head

	var cur *ListNode

	for head != nil {
		tmp := cur
		cur = &ListNode{Val: head.Val, Next: tmp}

		head = head.Next
	}

	return cur
}
