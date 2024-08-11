package main

import "fmt"

type Node[T comparable] struct {
	val  T
	next *Node[T]
}

// Appends a new node to the end of the list
func (root *Node[T]) Add(val T) *Node[T] {
	node := &Node[T]{val: val}
	if root == nil {
		return node
	}

	last := root.Last()
	last.next = node

	return root
}

// Returns the last node in the list, panics if the list is empty
func (root *Node[T]) Last() *Node[T] {
	if root == nil {
		panic("empty list")
	}
	for root.next != nil {
		root = root.next
	}
	return root
}

// Inserts a new node at the given index, panics if the index is out of range
func (root *Node[T]) Insert(val T, idx int) *Node[T] {
	node := &Node[T]{val: val}

	if idx == 0 {
		node.next = root
		return node
	}

	prev := root.Nth(idx - 1)
	node.next = prev.next
	prev.next = node

	return root
}

// Returns the nth node in the list, panics if the index is out of range
func (root *Node[T]) Nth(idx int) *Node[T] {
	if root == nil {
		panic("index out of range")
	}
	for i := 0; i < idx; i++ {
		root = root.next
		if root == nil {
			panic("index out of range")
		}
	}
	return root
}

// Returns the position of the supplied value, -1 if not found
func (root *Node[T]) Index(val T) int {
	for i := 0; root != nil; i++ {
		if root.val == val {
			return i
		}
		root = root.next
	}
	return -1
}

func main() {
	var root *Node[int]
	// []
	root = root.Add(1)
	// [1]
	root = root.Add(2)
	// [1,2]
	root = root.Add(3)
	// [1,2,3]
	root = root.Insert(4, 1)
	// [1,4,2,3]
	root = root.Insert(5, 0)
	// [5,1,4,2,3]
	root = root.Insert(6, 4)
	// [5,1,4,2,6,3]

	for node := root; node != nil; node = node.next {
		println(node.val)
	}

	fmt.Println(root.Index(1))  // 1
	fmt.Println(root.Index(10)) // -1
}
