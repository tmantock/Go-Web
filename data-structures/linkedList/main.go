package main

import "github.com/tmantock/Go-Web/data-structures/linkedList/singly"
import "fmt"

func main() {
	l := singly.List{}
	l.Append(5)
	l.Append(6)
	l.Append(2)
	l.Traverse()
	fmt.Println()
	l.Prepend(9)
	l.Traverse()
}
