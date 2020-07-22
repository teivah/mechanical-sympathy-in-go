package tests_test

import (
	"fmt"
	"testing"

	"github.com/clarketm/ncalc/hexadecimal"
)

const (
	it = 1_000_000_0
)

//type I1 struct {
//	b1 bool
//	i  int64
//	b2 bool
//}
//
//func BenchmarkI1(b *testing.B) {
//	s := make([]I1, it)
//	var r int64
//	b.ResetTimer()
//	for j := 0; j < it; j++ {
//		r += s[j].i
//	}
//	result = r
//}
//
//type I2 struct {
//	i  int64
//	b1 bool
//	b2 bool
//}
//
//func BenchmarkI2(b *testing.B) {
//	s := make([]I2, it)
//	var r int64
//	b.ResetTimer()
//	for j := 0; j < it; j++ {
//		r += s[j].i
//	}
//	result = r
//}
//
//const it = 100_000_0000
//
//var result int64
//
//func TestName(t *testing.T) {
//	b := I1{}
//
//	fmt.Println(unsafe.Sizeof(b))
//	fmt.Println(unsafe.Offsetof(b.b1))
//	fmt.Println(unsafe.Offsetof(b.i))
//	fmt.Println(unsafe.Offsetof(b.b2))
//
//	c := I2{}
//
//	fmt.Println(unsafe.Sizeof(c))
//	fmt.Println(unsafe.Offsetof(c.i))
//	fmt.Println(unsafe.Offsetof(c.b1))
//	fmt.Println(unsafe.Offsetof(c.b2))
//}

var result int32

type Struct1 struct {
	a int32
	b int64
}

func BenchmarkSliceOfStructures(b *testing.B) {
	s := make([]Struct1, it)
	var r int32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < it; i++ {
			r = s[i].a
		}
	}
	result = r
}

type Struct2 struct {
	a []int32
	b []int64
}

func BenchmarkStructureOfSlices(b *testing.B) {
	s := Struct2{
		a: make([]int32, it),
		b: make([]int64, it),
	}
	var r int32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < it; i++ {
			r = s.a[i]
		}
	}
	result = r
}

//type ElementOnly struct {
//	value int
//}
//
//func Benchmark_TraverseSliceOfElementOnly(b *testing.B) {
//	s := make([]ElementOnly, it)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		sum := 0
//		for i := 0; i < it; i++ {
//			sum += s[i].value
//		}
//	}
//}

type Struct struct {
}

func TestName(t *testing.T) {
	s := Struct{}
	s2 := Struct{}
	fmt.Printf("%v\n", hexadecimal.Hexadecimal2Decimal(fmt.Sprintf("%p", &s)))
	fmt.Printf("%v\n", hexadecimal.Hexadecimal2Decimal(fmt.Sprintf("%p", &s2)))
}
