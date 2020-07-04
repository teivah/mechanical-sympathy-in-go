package tests

import (
	"fmt"
	"testing"

	"github.com/clarketm/ncalc/hexadecimal"
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

//// False sharing example
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

//func BenchmarkLoopInterchange(b *testing.B) {
//	m := createMatrix(5000, 100)
//	runtime.GC()
//	b.ResetTimer()
//	for n := 0; n < b.N; n++ {
//		for j := 0; j < 100; j++ {
//			for i := 0; i < 5000; i++ {
//				m[i][j] = 2 * m[i][j]
//			}
//		}
//	}
//}
//
//func BenchmarkLoopInterchange2(b *testing.B) {
//	m := createMatrix(5000, 100)
//	runtime.GC()
//	b.ResetTimer()
//	for n := 0; n < b.N; n++ {
//		for i := 0; i < 5000; i++ {
//			for j := 0; j < 100; j++ {
//				m[i][j] = 2 * m[i][j]
//			}
//		}
//	}
//}

// Cache associativity example
//const size = 512
//
//var result int32
//
//func TestName(t *testing.T) {
//	const arraySize = 8
//	s := make([][arraySize]int64, size)
//	for j := 0; j < 10; j++ {
//		for i := 0; i < arraySize; i++ {
//			fmt.Printf("%d %d: %p \n", j, i, &s[j][i])
//		}
//	}
//}
//
//func Benchmark7Size(b *testing.B) {
//	const arraySize = 16
//	s := make([][arraySize]int32, size)
//	var r int32
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		for j := 0; j < size; j++ {
//			ss := s[j]
//			for k := 0; k < arraySize; k++ {
//				r = ss[k]
//			}
//		}
//	}
//	result = r
//}
//
//func Benchmark8Size(b *testing.B) {
//	const arraySize = 17
//	s := make([][arraySize]int32, size)
//	var r int32
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		for j := 0; j < size; j++ {
//			ss := s[j]
//			for k := 0; k < arraySize; k++ {
//				r = ss[k]
//			}
//		}
//	}
//	result = r
//}

type testx struct {
	data      [8]int64
	something [8192 - 64 + 8]byte
}

func TestName(t *testing.T) {
	const size = 1025
	s := make([][size]int64, size)
	//s := createMatrix(size, size)
	setDistribution := make(map[string]int)
	blockDistribution := make(map[string]map[string]int)
	for i := 0; i < size; i++ {
		v := parse(6, 7, &s[i])
		setDistribution[v.set]++
		if blockDistribution[v.set] == nil {
			blockDistribution[v.set] = make(map[string]int)
		}
		blockDistribution[v.set][v.tagBits]++
	}
	total := 0
	for _, v := range setDistribution {
		fmt.Printf("set: %v\n", v)
		total += v
		//for _, v2 := range blockDistribution[k] {
		//fmt.Printf("block %v\n", v2)
		//}
	}
	fmt.Printf("total: %v\n", total)
}

type info struct {
	set     string
	tagBits string
}

func parse(block, setSize int, i interface{}) info {
	b := hexadecimal.Hexadecimal2Binary(fmt.Sprintf("%p", i))
	fmt.Printf("%v\n", b)
	return info{
		set:     b[len(b)-block-setSize : len(b)-block],
		tagBits: b[:len(b)-block-setSize],
	}
}

//var result int64
//
//type bar struct {
//	data      [8]int64
//	something [1024]byte
//}
//
//func BenchmarkBar(b *testing.B) {
//	const arraySize = 1_000
//	s := [arraySize]bar{}
//	var r int64
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		for j := 0; j < arraySize; j++ {
//			ss := s[j].data
//			for k := 0; k < 8; k++ {
//				r = ss[k]
//			}
//		}
//	}
//	result = r
//}
//
//type foo struct {
//	data      [8]int64
//	something [512 - 64]byte
//}
//
//func BenchmarkFoo(b *testing.B) {
//	const arraySize = 1_000
//	s := [arraySize]foo{}
//	var r int64
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		for j := 0; j < arraySize; j++ {
//			ss := s[j].data
//			for k := 0; k < 8; k++ {
//				r = ss[k]
//			}
//		}
//	}
//	result = r
//}

func TestX(t *testing.T) {
	m := createMatrix(512, 512)
	transpose(m)
}

func BenchmarkTranspose1(b *testing.B) {
	const n = 2048
	m := createMatrix(n, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		transpose(m)
	}
}

func BenchmarkTranspose2(b *testing.B) {
	const n = 2049
	m := createMatrix(n, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		transpose(m)
	}
}

func transpose(m [][]int32) {
	length := len(m)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			m[i][j], m[j][i] = m[j][i], m[i][j]
		}
	}
}

// 100 11100 010000 // 10000
// 101 11100 000000 // 12032
// 110 11100 000000 // 14080
