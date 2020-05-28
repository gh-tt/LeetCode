package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := ListNode{4, &ListNode{5, nil}}
	re := swapPairs(&head)
	fmt.Println(re,re.Next)
}

func swapPairs(head *ListNode) *ListNode {

	dummy := new(ListNode)
	dummy.Next = head

	prev := dummy

	for head != nil && head.Next != nil {

		first := head
		second := head.Next

		prev.Next = second
		first.Next = second.Next
		second.Next = first

		prev = first
		head = first.Next
	}

	return dummy.Next
}
