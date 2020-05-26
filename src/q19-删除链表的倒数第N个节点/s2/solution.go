package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := ListNode{4, &ListNode{5, &ListNode{6, nil}}}
	re := removeNthFromEnd(&head, 1)
	fmt.Println(re, re.Next)
}

//双指针O(n)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur := new(ListNode)
	cur.Next = head

	first := cur
	second := cur

	for i := 1; i <= n+1; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next

	return cur.Next
}
