package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}
	fmt.Println(hasCycle(&head))
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	first := head
	second := head

	for first != nil && first.Next != nil {
		first = first.Next.Next
		second = second.Next
		if first == second {
			return true
		}
	}
	return false
}
