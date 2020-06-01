package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func main() {
	m := make(map[string]int)

	m["q"] = 1
	fmt.Println(m["w"])
}

func printList(l *Node) {
	for l != nil {
		fmt.Printf("%+v  %p\n", l, l)
		l = l.Next
	}
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	q, newHead := head, &Node{}

	p := newHead

	m := make(map[*Node]*Node)

	for q != nil {

		if m[q] == nil {
			m[q] = &Node{Val: q.Val}
		}

		p.Next = m[q]
		p = p.Next

		if q.Random == nil {
			m[q].Random = nil
		} else {
			if m[q].Random == nil {
				m[q.Random] = &Node{Val: q.Random.Val}
			}
			p.Random = m[q.Random]
		}


		q = q.Next
	}

	return newHead.Next
}
