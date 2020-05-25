package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := ListNode{4, &ListNode{5, &ListNode{6,nil}}}
	re := removeNthFromEnd(&head, 2)
	fmt.Println(re,re.Next)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	a := new(ListNode)

	a.Next = head

	first := head

	length := 0
	for first != nil {
		length++
		first = first.Next
	}

	length -= n
	first = a
	for length > 0 {
		length--
		first = first.Next
	}

	first.Next = first.Next.Next

	return a.Next
}
