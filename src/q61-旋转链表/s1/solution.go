package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}
	//printList(&head)
	printList(rotateRight(&head,11))
}
func printList(l *ListNode) {
	for l != nil {
		fmt.Printf("%+v  %p\n", l, l)
		l = l.Next
	}
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	dummy := new(ListNode)

	dummy.Next = head

	first := dummy
	begin := head

	end :=new(ListNode)
	length := 0
	for head != nil {
		length++
		end = head
		head = head.Next
	}

	mod := k % length
	if mod == 0 {
		return dummy.Next
	}

	for i := 1; i <= length-mod; i++ {
		first = first.Next
	}

	dummy.Next = first.Next
	first.Next = nil
	end.Next = begin

	return dummy.Next
}
