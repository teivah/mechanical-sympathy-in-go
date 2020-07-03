package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	iterations, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", func2(iterations))
}

type bar struct {
	data      [8]int64
	something [8192 - 64]byte
}

func func2(n int) int64 {
	const size = 512
	s := make([]bar, size)
	var r int64
	for i := 0; i < n; i++ {
		for j := 0; j < size; j++ {
			cacheLine := s[j].data
			for k := 0; k < 8; k++ {
				r = cacheLine[k]
			}
		}
	}
	return r
}
