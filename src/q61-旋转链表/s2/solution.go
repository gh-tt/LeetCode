package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}
	//printList(&head)
	printList(rotateRight(&head, 8))
}
func printList(l *ListNode) {
	for l != nil {
		fmt.Printf("%+v  %p\n", l, l)
		l = l.Next
	}
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	cur := head
	n := 1
	for cur.Next != nil {
		n++
		cur = cur.Next
	}

	cur.Next = head

	mod := k % n
	if mod == 0 {
		cur.Next = nil
		return head
	}
	step := n - mod

	for step > 0 {
		cur = cur.Next
		step--
	}

	newHead := cur.Next
	cur.Next = nil

	return newHead
}
