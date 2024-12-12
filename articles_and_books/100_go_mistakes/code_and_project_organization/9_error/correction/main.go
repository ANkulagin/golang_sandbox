package main

import "fmt"

type Node[T any] struct {
	Val  T
	Next *Node[T]
}

func (n *Node[T]) Add(next *Node[T]) {
	current := n
	for current.Next != nil {
		current = current.Next
	}
	current.Next = next
}

func main() {
	head := &Node[int]{Val: 1}
	head.Add(&Node[int]{Val: 2})
	head.Add(&Node[int]{Val: 3})

	current := head
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}
