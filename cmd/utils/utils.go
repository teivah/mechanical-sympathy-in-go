package utils

const (
	defaultValue = 1
)

func CreateSlice(n int) []int {
	return make([]int, n)
}

func CreateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	return matrix
}

func CreateSliceOfStructures(n int) []Element {
	s := make([]Element, n)
	for i := 0; i < n; i++ {
		s[i] = Element{}
	}
	return s
}

type Element struct {
	Value     int
	Something [1024]byte
}

func CreateStructureOfSlices(n int) Element2 {
	return Element2{
		Values:    make([]int, n),
		Something: make([][1024]byte, n),
	}
}

type Element2 struct {
	Values    []int
	Something [][1024]byte
}
