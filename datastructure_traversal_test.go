package tests

import (
	"testing"
)

const (
	defaultValue = 1
	iteration    = 1_000_000
)

func Benchmark_TraverseSlice(b *testing.B) {
	s := createSlice(iteration)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < iteration; i++ {
			sum += s[i]
		}
	}
}

func createSlice(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = defaultValue
	}
	return s
}

func Benchmark_TraverseLinkedList_BoundSize(b *testing.B) {
	root := createLinkedList(iteration)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		node := root
		sum := 0
		for i := 0; i < iteration; i++ {
			sum += node.value
			node = node.next
		}
	}
}

func Benchmark_TraverseLinkedList_Nullity(b *testing.B) {
	root := createLinkedList(iteration)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		node := root
		sum := 0
		for node != nil {
			sum += node.value
			node = node.next
		}
	}
}

func createLinkedList(n int) *Node {
	root := &Node{value: defaultValue}
	current := root

	for i := 0; i < n; i++ {
		node := &Node{value: defaultValue}
		current.next = node
		current = node
	}
	return root
}

type Node struct {
	value int
	next  *Node
}

func Benchmark_TraverseStridingStructure(b *testing.B) {
	s := createStridingStructure(iteration)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < iteration; i++ {
			sum += s[i].value
		}
	}
}

func createStridingStructure(n int) []Element {
	s := make([]Element, n)
	for i := 0; i < n; i++ {
		s[i] = Element{
			value: defaultValue,
		}
	}
	return s
}

type Element struct {
	value int
}
