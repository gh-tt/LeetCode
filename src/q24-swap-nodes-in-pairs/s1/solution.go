package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := ListNode{4, &ListNode{5, nil}}
	printList(&head)
	newList := swapPairs(&head)
	printList(newList)

	newList = swapPairs(newList)
	printList(newList)
	//re := swapPairs(&head)

	//printList(&head)
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Printf("%+v  %p\n", l, l)
		l = l.Next
	}
}

func swapPairs(head *ListNode) *ListNode {

	dummy := new(ListNode)
	dummy.Next = head

	prev := dummy

	for head != nil && head.Next != nil {

		first := head.Next

		prev.Next = first
		head.Next = first.Next
		first.Next = head

		prev = head
		head = head.Next
	}

	return dummy.Next
}
