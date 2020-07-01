package tests

import (
	"sync"
	"testing"
)

//func Benchmark_TraverseSliceOfInts(b *testing.B) {
//	s := createSlice(iteration)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		sum := 0
//		for i := 0; i < iteration; i++ {
//			sum += s[i]
//		}
//	}
//}
//
//func Benchmark_TraverseSliceOfInts_Reverse(b *testing.B) {
//	s := createSlice(iteration)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		sum := 0
//		for i := iteration - 1; i >= 0; i-- {
//			sum += s[i]
//		}
//	}
//}

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
//func Benchmark_TraverseSliceOfElementOnly(b *testing.B) {
//	s := createElementOnly(iteration)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		sum := 0
//		for i := 0; i < iteration; i++ {
//			sum += s[i].value
//		}
//	}
//}
//
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

func BenchmarkStructureFalseSharing(b *testing.B) {
	structA := SimpleStruct{}
	structB := SimpleStruct{}
	wg := sync.WaitGroup{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(2)
		go func() {
			for j := 0; j < iteration; j++ {
				structA.n += j
			}
			wg.Done()
		}()
		go func() {
			for j := 0; j < iteration; j++ {
				structB.n += j
			}
			wg.Done()
		}()
		wg.Wait()
	}
	b.StopTimer()
}

func BenchmarkStructurePadding(b *testing.B) {
	structA := PaddedStruct{}
	structB := SimpleStruct{}
	wg := sync.WaitGroup{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(2)
		go func() {
			for j := 0; j < iteration; j++ {
				structA.n += j
			}
			wg.Done()
		}()
		go func() {
			for j := 0; j < iteration; j++ {
				structB.n += j
			}
			wg.Done()
		}()
		wg.Wait()
	}
	b.StopTimer()
}
