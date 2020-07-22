package main

import (
	"fmt"
)

type Element struct {
	Value     int
	Something [1024]byte
}

func main() {
	const iterations = 100
	s := make([]Element, iterations)
	sum := 0

	for i := 0; i < iterations; i++ {
		sum += s[i].Value
	}

	fmt.Printf("%v\n", sum)

	for i := 0; i < 1000; i++ {
		sum += s[i].Value
	}
}
