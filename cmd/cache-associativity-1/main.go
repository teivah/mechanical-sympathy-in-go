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

func func1(n int) int64 {
	//const size = 1024 // 7.48
	//const size = 512 // 6.94
	const size = 1025 // 4.28
	//const size = 256 // 6.62
	//const size = 128 // 6.17
	//const size = 64 // 5.21
	s := make([][size]int64, n)
	var r int64
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < 8; k++ {
				r = s[j][k]
			}
		}
	}
	return r
}
