package tests

//const (
//	it = 1_000_000
//)
//
//var result int64
//
//func BenchmarkForwardsIteration(b *testing.B) {
//	s := make([]int64, it)
//	var r int64
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		length := len(s)
//		for i := 0; i < length; i++ {
//			r = s[i]
//		}
//	}
//	result = r
//}
//
//func BenchmarkBackwardsIteration(b *testing.B) {
//	s := make([]int64, it)
//	var r int64
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		length := len(s)
//		for i := length - 1; i >= 0; i-- {
//			r = s[i]
//		}
//	}
//	result = r
//}
