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

func func2(n int) int64 {
	const size = 1025
	s := make([][size]int64, n)
	var r int64
	for i := 0; i < n; i++ {
		for j := 0; j < 1024; j++ {
			for k := 0; k < 8; k++ {
				r = s[j][k]
			}
		}
	}
	return r
}
