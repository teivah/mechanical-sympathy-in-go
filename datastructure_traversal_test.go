package tests

import "testing"

const (
	defaultValue     = 1
	iteration        = 1_000_000
	matrix           = 1_000
	CacheLinePadSize = 64
)

func Benchmark_TraverseSliceOfInts(b *testing.B) {
	s := createSlice(iteration)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < iteration; i++ {
			sum += s[i]
		}
	}
}

func Benchmark_TraverseSliceOfInts_Reverse(b *testing.B) {
	s := createSlice(iteration)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum := 0
		for i := iteration - 1; i >= 0; i-- {
			sum += s[i]
		}
	}
}

func Benchmark_TraverseSliceOfInts_Add(b *testing.B) {
	s := createSlice(iteration)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 0; i < iteration; {
			sum += s[i]
			i += s[i]
		}
	}
}

func Benchmark_TraverseSliceOfInts_Reverse_Add(b *testing.B) {
	s := createSlice(iteration)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum := 0
		for i := iteration - 1; i >= 0; {
			sum += s[i]
			i -= s[i]
		}
	}
}

//func Benchmark_TraverseLinkedList_BoundSize(b *testing.B) {
//	root := createLinkedList(iteration)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		node := root
//		sum := 0
//		for i := 0; i < iteration; i++ {
//			sum += node.value
//			node = node.next
//		}
//	}
//}
//
//func Benchmark_TraverseLinkedList_Nullity(b *testing.B) {
//	root := createLinkedList(iteration)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		node := root
//		sum := 0
//		for node != nil {
//			sum += node.value
//			node = node.next
//		}
//	}
//}

//// Constant stride
//func Benchmark_TraverseSliceOfStructures(b *testing.B) {
//	s := createSliceOfStructures(iteration)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		sum := 0
//		for i := 0; i < iteration; i++ {
//			sum += s[i].value
//		}
//	}
//}
//
//// Unit stride
//func Benchmark_TraverseStructureOfSlices(b *testing.B) {
//	s := createStructureOfSlices(iteration)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		sum := 0
//		for i := 0; i < iteration; i++ {
//			sum += s.values[i]
//		}
//	}
//}

//func Benchmark_TraverseMatrix_Row(b *testing.B) {
//	m := createMatrix(matrix)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		sum := 0
//		for i := 0; i < matrix; i++ {
//			for j := 0; j < matrix; j++ {
//				sum += m[i][j]
//			}
//		}
//	}
//}
//
//func Benchmark_TraverseMatrix_Column(b *testing.B) {
//	m := createMatrix(matrix)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		sum := 0
//		for i := 0; i < matrix; i++ {
//			for j := 0; j < matrix; j++ {
//				sum += m[j][i]
//			}
//		}
//	}
//}

func createSlice(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = defaultValue
	}
	return s
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

func createSliceOfStructures(n int) []Element {
	s := make([]Element, n)
	for i := 0; i < n; i++ {
		s[i] = Element{
			value: defaultValue,
		}
	}
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

	for i := 0; i < n; i++ {
		elem.values[i] = defaultValue
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
		for j := 0; j < n; j++ {
			matrix[i][j] = defaultValue
		}
	}
	return matrix
}
