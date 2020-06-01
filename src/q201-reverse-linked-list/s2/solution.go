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
	var prev *ListNode = nil

	cur := head

	for cur != nil {
		tmp := cur.Next

		cur.Next = prev

		prev = cur

		cur = tmp

	}
	return prev
}
