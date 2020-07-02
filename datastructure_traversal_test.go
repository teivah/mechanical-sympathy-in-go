package tests

import (
	"runtime"
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

//func BenchmarkStructureFalseSharing(b *testing.B) {
//	structA := SimpleStruct{}
//	structB := SimpleStruct{}
//	wg := sync.WaitGroup{}
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		wg.Add(2)
//		go func() {
//			for j := 0; j < iteration; j++ {
//				structA.n += j
//			}
//			wg.Done()
//		}()
//		go func() {
//			for j := 0; j < iteration; j++ {
//				structB.n += j
//			}
//			wg.Done()
//		}()
//		wg.Wait()
//	}
//	b.StopTimer()
//}
//
//func BenchmarkStructurePadding(b *testing.B) {
//	structA := PaddedStruct{}
//	structB := SimpleStruct{}
//	wg := sync.WaitGroup{}
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		wg.Add(2)
//		go func() {
//			for j := 0; j < iteration; j++ {
//				structA.n += j
//			}
//			wg.Done()
//		}()
//		go func() {
//			for j := 0; j < iteration; j++ {
//				structB.n += j
//			}
//			wg.Done()
//		}()
//		wg.Wait()
//	}
//	b.StopTimer()
//}

// False sharing example
//var result int
//
//func BenchmarkStructureFalseSharing(b *testing.B) {
//	s := createSlice(iterationFalseSharing)
//	var counter Result
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		counter = countOdds(s)
//	}
//	b.StopTimer()
//	result = counter.odds
//}
//
//type Result struct {
//	odds int
//}
//
//func countOdds(s []int) Result {
//	counter := Result{}
//	for i := 0; i < len(s); i++ {
//		if s[i]%2 == 0 {
//			counter.odds++
//		}
//	}
//	return counter
//}
//
//func BenchmarkStructureFalseSharingConcurrent(b *testing.B) {
//	s := createSlice(iterationFalseSharing)
//	var counter ConcurrentResult
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		counter = countOddsButConcurrent(s)
//	}
//	b.StopTimer()
//	result = counter.firstHalfOdds + counter.secondHalfOdds
//}
//
//type ConcurrentResult struct {
//	firstHalfOdds  int
//	secondHalfOdds int
//}
//
//func countOddsButConcurrent(s []int) ConcurrentResult {
//	counter := ConcurrentResult{}
//	wg := sync.WaitGroup{}
//	wg.Add(2)
//	go func() {
//		for i := 0; i < len(s)/2; i++ {
//			if s[i]%2 == 0 {
//				counter.firstHalfOdds++
//			}
//		}
//		wg.Done()
//	}()
//	go func() {
//		for i := len(s) / 2; i < len(s); i++ {
//			if s[i]%2 == 0 {
//				counter.secondHalfOdds++
//			}
//		}
//		wg.Done()
//	}()
//	wg.Wait()
//	return counter
//}
//
//func BenchmarkStructureFalseSharingConcurrent2(b *testing.B) {
//	s := createSlice(iterationFalseSharing)
//	var counter ConcurrentResultWithPadding
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		counter = countOddsButConcurrent2(s)
//	}
//	b.StopTimer()
//	result = counter.firstHalfOdds + counter.secondHalfOdds
//}
//
//type ConcurrentResultWithPadding struct {
//	firstHalfOdds  int
//	_              CacheLinePad
//	secondHalfOdds int
//}
//
//func countOddsButConcurrent2(s []int) ConcurrentResultWithPadding {
//	counter := ConcurrentResultWithPadding{}
//	wg := sync.WaitGroup{}
//	wg.Add(2)
//	go func() {
//		for i := 0; i < len(s)/2; i++ {
//			if s[i]%2 == 0 {
//				counter.firstHalfOdds++
//			}
//		}
//		wg.Done()
//	}()
//	go func() {
//		for i := len(s) / 2; i < len(s); i++ {
//			if s[i]%2 == 0 {
//				counter.secondHalfOdds++
//			}
//		}
//		wg.Done()
//	}()
//	wg.Wait()
//	return counter
//}

func BenchmarkLoopInterchange(b *testing.B) {
	m := createMatrix(5000, 100)
	runtime.GC()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for j := 0; j < 100; j++ {
			for i := 0; i < 5000; i++ {
				m[i][j] = 2 * m[i][j]
			}
		}
	}
}

func BenchmarkLoopInterchange2(b *testing.B) {
	m := createMatrix(5000, 100)
	runtime.GC()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < 5000; i++ {
			for j := 0; j < 100; j++ {
				m[i][j] = 2 * m[i][j]
			}
		}
	}
}
