package tests

const (
	iteration             = 1_000_000
	iterationFalseSharing = 500_000_000
	matrix                = 1_000
	CacheLinePadSize      = 64
)

type CacheLinePad struct {
	_ [CacheLinePadSize]byte
}

func createSlice(n int) []int {
	s := make([]int, n)
	return s
}

func createMatrix(n, m int) [][]int32 {
	matrix := make([][]int32, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int32, m)
	}
	return matrix
}
