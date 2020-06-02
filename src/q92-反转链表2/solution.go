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

	p := dummy
	q := head

	for i := 0; i < m-1; i++ {
		p = p.Next
		q = q.Next
	}

	s, e := p, q

	for j := 0; j <= n-m; j++ {
		tmp := q.Next

		q.Next = p

		p = q

		q = tmp
	}
	s.Next = p
	e.Next = q
	return dummy.Next

}
