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
