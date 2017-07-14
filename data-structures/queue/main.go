package main

import (
	"fmt"
)

type List interface {
	Push(n *Node)
	Pop() (n *Node)
	Peek() (n *Node)
}

func Peek(l List) (n *Node) {
	return l.Peek()
}

func main() {
	var q Queue

	n := &Node{5}
	f := &Node{4}

	q.Push(n)
	q.Push(f)
	p := Peek(&q)
	fmt.Println(p)
	fmt.Println(q)

	q.Pop()
	q.Pop()
	fmt.Println(q)
}
