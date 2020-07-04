package main

import (
	"log"
	"os"
	"strconv"

	"github.com/teivah/mechanical-sympathy-in-go/cmd/utils"
)

func main() {
	iterations, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	m := utils.CreateMatrix(2049)
	for i := 0; i < iterations; i++ {
		transpose(m)
	}
}

func transpose(m [][]int32) {
	length := len(m)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			m[i][j], m[j][i] = m[j][i], m[i][j]
		}
	}
}
