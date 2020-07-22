package tests

import (
	"fmt"
	"testing"
	"time"

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
//	structA := Struct{}
//	structB := Struct{}
//	wg := sync.WaitGroup{}
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		wg.Add(2)
//		go func() {
//			for j := 0; j < iteration; j++ {
//				structA.var1 += j
//			}
//			wg.Done()
//		}()
//		go func() {
//			for j := 0; j < iteration; j++ {
//				structB.var1 += j
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
//	structB := Struct{}
//	wg := sync.WaitGroup{}
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		wg.Add(2)
//		go func() {
//			for j := 0; j < iteration; j++ {
//				structA.var1 += j
//			}
//			wg.Done()
//		}()
//		go func() {
//			for j := 0; j < iteration; j++ {
//				structB.var1 += j
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
//	for var1 := 0; var1 < b.N; var1++ {
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
//	for var1 := 0; var1 < b.N; var1++ {
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
//			fmt.Printf("%d %d: %p \var1", j, i, &s[j][i])
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
	const size = 1024
	s := make([][size]byte, size)

	for i := 0; i < 8; i++ {
		parse(6, 7, &s[0][i])
	}
	parse(2, 7, &s[1][0])

	//setDistribution := make(map[string]int)
	//blockDistribution := make(map[string]map[string]int)
	//for i := 0; i < size; i++ {
	//	v := parse(6, 7, &s[i])
	//	setDistribution[v.set]++
	//	if blockDistribution[v.set] == nil {
	//		blockDistribution[v.set] = make(map[string]int)
	//	}
	//	blockDistribution[v.set][v.tagBits]++
	//}
	//total := 0
	//for _, v := range setDistribution {
	//	fmt.Printf("set: %v\var1", v)
	//	total += v
	//	//for _, v2 := range blockDistribution[k] {
	//	//fmt.Printf("block %v\var1", v2)
	//	//}
	//}
	//fmt.Printf("total: %v\var1", total)
}

func TestName2(t *testing.T) {
	const size = 1024

	go func() {
		s := make([][size]byte, size)
		i := parse(6, 7, &s[0]) // 1100000000000000000011000010000000000000 824634515456
		fmt.Printf("%#v\n", i.binary)
	}()
	go func() {
		s := make([][size]byte, size)
		i := parse(6, 7, &s[0]) // 1100000000000000000111001100000000000000 824635604992
		fmt.Printf("%#v\n", i.binary)
	}()
	time.Sleep(250 * time.Millisecond)

}

type info struct {
	set     string
	tagBits string
	binary  string
}

func parse(block, setSize int, i interface{}) info {
	b := hexadecimal.Hexadecimal2Binary(fmt.Sprintf("%p", i))
	return info{
		set:     b[len(b)-block-setSize : len(b)-block],
		tagBits: b[:len(b)-block-setSize],
		binary:  b,
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

//func BenchmarkTranspose1(b *testing.B) {
//	const n = 2048
//	m := createMatrix(n, n)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		transpose(m)
//	}
//}
//
//func BenchmarkTranspose2(b *testing.B) {
//	const n = 2049
//	m := createMatrix(n, n)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		transpose(m)
//	}
//}

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
