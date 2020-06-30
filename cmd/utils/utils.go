package utils

const (
	defaultValue = 1
)

func CreateSlice(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = defaultValue
	}
	return s
}
