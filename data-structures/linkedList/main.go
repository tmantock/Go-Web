package main

import (
	"fmt"

	"github.com/tmantock/Go-Web/data-structures/linkedList/doubly"
	"github.com/tmantock/Go-Web/data-structures/linkedList/singly"
)

type LinkedList interface {
	Traverse()
}

func Traverse(l LinkedList) {
	l.Traverse()
}

func main() {
	l := singly.List{}
	l.Append(5)
	l.Append(6)
	l.Append(2)
	Traverse(l)
	fmt.Println()
	l.Prepend(9)
	Traverse(l)
	fmt.Println()
	ll := doubly.List{}
	ll.Append(3)
	ll.Append(7)
	Traverse(ll)
}
