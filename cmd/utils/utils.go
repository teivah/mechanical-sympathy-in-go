package utils

const (
	defaultValue = 1
	Iteration    = 10_000_000
)

func CreateSlice(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = defaultValue
	}
	return s
}
