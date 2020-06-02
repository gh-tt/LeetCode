package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}
	printList(reverseBetween(&head, 2, 3))
}
func printList(l *ListNode) {
	for l != nil {
		fmt.Printf("%+v  %p\n", l, l)
		l = l.Next
	}
}

func reverseBetween(head *ListNode, m, n int) *ListNode {
	dummy := new(ListNode)

	dummy.Next = head

	prev := dummy

	for i := 0; i < m-1; i++ {
		prev = prev.Next
	}
	cur := prev.Next
	for j := m; j < n; j++ {
		tmp := cur.Next

		cur.Next = tmp.Next
		tmp.Next = prev.Next
		prev.Next = tmp
	}
	return dummy.Next
}
