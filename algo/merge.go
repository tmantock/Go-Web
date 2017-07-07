package main

import (
	"fmt"
)

func main() {
	s := []int{5, 3, 6, 8, 2, 7, 1, 9, 4}
	x := sort(s)
	fmt.Println(x)
}

func sort(s []int) []int {
	if len(s) < 2 {
		return s[:]
	}

	mid := len(s) / 2
	var left []int
	var right []int
	left = sort(s[:mid])
	right = sort(s[mid:len(s)])
	return merge(left, right)
}

func merge(l []int, r []int) []int {
	i := 0
	j := 0
	var sorted []int
	for i < len(l) && j < len(r) {
		if l[i] < r[j] {
			sorted = append(sorted, l[i])
			i++
		} else {
			sorted = append(sorted, r[j])
			j++
		}
	}

	for i < len(l) {
		sorted = append(sorted, l[i])
		i++
	}

	for j < len(r) {
		sorted = append(sorted, r[j])
		j++
	}

	return sorted
}
