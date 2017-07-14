package main

type Stack []*Node

func (s *Stack) Push(n *Node) {
	*s = append(*s, n)
}

func (s *Stack) Pop() (n *Node) {
	x := s.Len() - 1
	n = (*s)[x]
	*s = (*s)[:x]
	return
}

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Peek() (n *Node) {
	return (*s)[len(*s)-1]
}
