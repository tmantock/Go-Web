package doubly

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) Append(data int) {
	n := &Node{
		data: data,
		prev: nil,
		next: nil,
	}

	if l.head == nil {
		l.head = n
		return
	}

	if l.tail == nil {
		l.tail = n
		l.head.next = l.tail
		l.tail.prev = l.head
		return
	}

	l.tail.next = n
	n.prev = l.tail
	l.tail = n
}

func (l List) Traverse() {
	if l.head == nil {
		return
	}

	root := l.head

	for root != nil {
		fmt.Println(root.data)
		root = root.next
	}
}
