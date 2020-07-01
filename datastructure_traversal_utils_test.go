package tests

const (
	iteration        = 1_000_000
	matrix           = 1_000
	CacheLinePadSize = 64
)

type CacheLinePad struct {
	_ [CacheLinePadSize]byte
}

func createSlice(n int) []int {
	s := make([]int, n)
	return s
}

func createLinkedList(n int) *Node {
	root := &Node{}
	current := root

	for i := 0; i < n; i++ {
		node := &Node{}
		current.next = node
		current = node
	}
	return root
}

type Node struct {
	value int
	next  *Node
}

func createElementOnly(n int) []ElementOnly {
	s := make([]ElementOnly, n)
	return s
}

type ElementOnly struct {
	value int
}

func createSliceOfStructures(n int) []Element {
	s := make([]Element, n)
	return s
}

type Element struct {
	value     int
	something [1024]byte
}

func createStructureOfSlices(n int) Element2 {
	elem := Element2{
		values:    make([]int, n),
		something: make([][1024]byte, n),
	}
	return elem
}

type Element2 struct {
	values    []int
	something [][1024]byte
}

func createMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	return matrix
}

type SimpleStruct struct {
	n int
}

type PaddedStruct struct {
	n int
	_ CacheLinePad
}
