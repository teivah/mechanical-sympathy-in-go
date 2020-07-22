package tests

//type Struct struct {
//	n int
//}
//
//var result int
//
//func BenchmarkIteration(b *testing.B) {
//	structA := Struct{} // Initialization
//	structB := Struct{} // Initialization
//	wg := sync.WaitGroup{}
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		wg.Add(2)
//		go func() { // Spin up first goroutine
//			for j := 0; j < iteration; j++ {
//				structA.n += j
//			}
//			wg.Done()
//		}()
//		go func() { // Spin up second goroutine
//			for j := 0; j < iteration; j++ {
//				structB.n += j
//			}
//			wg.Done()
//		}()
//		wg.Wait()                      // Wait
//		result = structA.n + structB.n // Aggregate
//	}
//}
//
//func BenchmarkIterationCommunication(b *testing.B) {
//	ch := make(chan int, 2)
//	for i := 0; i < b.N; i++ {
//		go func() { // Spin up first goroutine
//			i := 0 // Local state
//			for j := 0; j < iteration; j++ {
//				i += j
//			}
//			ch <- i
//		}()
//		go func() { // Spin up second goroutine
//			i := 0 // Local state
//			for j := 0; j < iteration; j++ {
//				i += j
//			}
//			ch <- i
//		}()
//		result = <-ch + <-ch // Wait and aggregate
//	}
//}
//
//type PaddedStruct struct {
//	_ cpu.CacheLinePad
//	n int
//	_ cpu.CacheLinePad
//}
//
//func BenchmarkIterationWithPadding(b *testing.B) {
//	structA := PaddedStruct{} // Initialization
//	structB := PaddedStruct{} // Initialization
//	wg := sync.WaitGroup{}
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		wg.Add(2)
//		go func() { // Spin up first goroutine
//			for j := 0; j < iteration; j++ {
//				structA.n += j
//			}
//			wg.Done()
//		}()
//		go func() { // Spin up second goroutine
//			for j := 0; j < iteration; j++ {
//				structB.n += j
//			}
//			wg.Done()
//		}()
//		wg.Wait() // Wait
//	}
//}

//// False sharing example
//var result int
//
//func BenchmarkSequential(b *testing.B) {
//	s := make([]int, iterationFalseSharing)
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
//func BenchmarkFalseSharingConcurrent(b *testing.B) {
//	s := make([]int, iterationFalseSharing)
//	var counter ConcurrentResult
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		counter = countOddsConcurrent(s)
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
//func countOddsConcurrent(s []int) ConcurrentResult {
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
//func BenchmarkFalseSharingConcurrentPadding(b *testing.B) {
//	s := make([]int, iterationFalseSharing)
//	var counter ConcurrentResultWithPadding
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		counter = countOddsConcurrentPadding(s)
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
//func countOddsConcurrentPadding(s []int) ConcurrentResultWithPadding {
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
//
//func BenchmarkFalseSharingConcurrentPadding2(b *testing.B) {
//	s := make([]int, iterationFalseSharing)
//	var counter ConcurrentResult
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		counter = countOddsConcurrentPadding2(s)
//	}
//	b.StopTimer()
//	result = counter.firstHalfOdds + counter.secondHalfOdds
//}
//
//type ConcurrentResultWithPadding2 struct {
//	firstHalfOdds int
//	_             CacheLinePad
//}
//
//func countOddsConcurrentPadding2(s []int) ConcurrentResult {
//	counterA := ConcurrentResultWithPadding{}
//	counterB := ConcurrentResultWithPadding{}
//	wg := sync.WaitGroup{}
//	wg.Add(2)
//	go func() {
//		for i := 0; i < len(s)/2; i++ {
//			if s[i]%2 == 0 {
//				counterA.firstHalfOdds++
//			}
//		}
//		wg.Done()
//	}()
//	go func() {
//		for i := len(s) / 2; i < len(s); i++ {
//			if s[i]%2 == 0 {
//				counterB.secondHalfOdds++
//			}
//		}
//		wg.Done()
//	}()
//	wg.Wait()
//	return ConcurrentResult{
//		firstHalfOdds:  counterA.firstHalfOdds,
//		secondHalfOdds: counterB.secondHalfOdds,
//	}
//}
//
//func TestCPU(t *testing.T) {
//	_ = cpu.CacheLinePad{}
//	x86 := cpu.X86
//	fmt.Printf("%#v\n", x86)
//}
