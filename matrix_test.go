package tests

import "testing"

const matrixLength = 6400

func BenchmarkMatrixCombination(b *testing.B) {
	matrixA := createInt64Matrix(matrixLength)
	matrixB := createInt64Matrix(matrixLength)

	for n := 0; n < b.N; n++ {
		for i := 0; i < matrixLength; i++ {
			for j := 0; j < matrixLength; j++ {
				matrixA[i][j] = matrixA[i][j] + matrixB[i][j]
			}
		}
	}
}

func BenchmarkMatrixReversedCombination(b *testing.B) {
	matrixA := createInt64Matrix(matrixLength)
	matrixB := createInt64Matrix(matrixLength)

	for n := 0; n < b.N; n++ {
		for i := 0; i < matrixLength; i++ {
			for j := 0; j < matrixLength; j++ {
				matrixA[i][j] = matrixA[i][j] + matrixB[j][i]
			}
		}
	}
}

func BenchmarkMatrixReversedCombinationPerBlock(b *testing.B) {
	matrixA := createInt64Matrix(matrixLength)
	matrixB := createInt64Matrix(matrixLength)

	blockSize := 8
	for n := 0; n < b.N; n++ {
		for i := 0; i < matrixLength; i += blockSize {
			for j := 0; j < matrixLength; j += blockSize {
				for ii := i; ii < i+blockSize; ii++ {
					for jj := j; jj < j+blockSize; jj++ {
						matrixA[ii][jj] = matrixA[ii][jj] + matrixB[jj][ii]
					}
				}
			}
		}
	}
}

func BenchmarkMatrixReversedCombinationPerBlock3(b *testing.B) {
	matrixA := createInt64Matrix(matrixLength)
	matrixB := createInt64Matrix(matrixLength)
	blockSize := 64

	for n := 0; n < b.N; n++ {
		for i := 0; i < matrixLength; i += blockSize {
			for j := 0; j < matrixLength; j += blockSize {
				for ii := i; ii < i+blockSize; ii++ {
					for jj := j; jj < j+blockSize; jj++ {
						matrixA[ii][jj] = matrixA[ii][jj] + matrixB[jj][ii]
					}
				}
			}
		}
	}
}

//func BenchmarkMatrixReversedCombinationPerBlock2(b *testing.B) {
//	matrixA := createInt64Matrix(matrixLength)
//	matrixB := createInt64Matrix(matrixLength)
//	blockSize := 4
//
//	for n := 0; n < b.N; n++ {
//		for i := 0; i < matrixLength; i += blockSize {
//			for j := 0; j < matrixLength; j += blockSize {
//				for ii := i; ii < i+blockSize; ii++ {
//					for jj := j; jj < j+blockSize; jj++ {
//						matrixA[ii][jj] = matrixA[ii][jj] + matrixB[jj][ii]
//					}
//				}
//			}
//		}
//	}
//}

func createInt64Matrix(size int) [][]int64 {
	matrix := make([][]int64, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int64, size)
		for j := 0; j < size; j++ {
			matrix[i][j] = 2
		}
	}
	return matrix
}
