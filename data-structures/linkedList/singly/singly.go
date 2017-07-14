package singly

import "fmt"

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) Append(data int) {
	node := &Node{
		data: data,
		next: nil,
	}
	if l.head == nil {
		l.head = node
		return
	}

	if l.tail == nil {
		l.head.next = node
		l.tail = node
		return
	}

	l.tail.next = node
	l.tail = node
}

func (l *List) Prepend(data int) {
	node := &Node{
		data: data,
		next: nil,
	}

	if l.head == nil {
		l.head = node
		return
	}

	node.next = l.head
	l.head = node
}

func (l *List) Traverse() {
	if l.head == nil {
		return
	}

	root := l.head

	for root != nil {
		fmt.Println(root.data)
		root = root.next
	}
}
