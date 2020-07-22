package tests

import (
	"testing"
)

const (
	it = 1_000_000
)

var result int64

func BenchmarkSliceIteration(b *testing.B) {
	s := make([]int64, it)
	var r int64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		length := len(s)
		for i := 0; i < length; i++ {
			r = s[i]
		}
	}
	result = r
}

type Node struct {
	Value int64
	Next  *Node
}

func CreateLinkedList(n int) *Node {
	root := &Node{}
	current := root

	for i := 0; i < n; i++ {
		node := &Node{}
		current.Next = node
		current = node
	}
	return root
}

func BenchmarkLinkedListIteration(b *testing.B) {
	root := CreateLinkedList(it)
	var r int64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		current := root
		for current != nil {
			r = current.Value
			current = current.Next
		}
	}
	result = r
}
