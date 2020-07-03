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
	fmt.Printf("%v\n", func1(iterations))
}

type foo struct {
	data      [8]int64
	something [16384 - 64]byte
}

func func1(n int) int64 {
	const size = 512
	s := make([]foo, size)
	var r int64
	for i := 0; i < n; i++ {
		for j := 0; j < size; j++ {
			ss := s[j].data
			for k := 0; k < 8; k++ {
				r = ss[k]
			}
		}
	}
	return r
}
