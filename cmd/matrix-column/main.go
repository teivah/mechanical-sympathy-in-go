package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/teivah/mechanical-sympathy-in-go/cmd/utils"
)

func main() {
	in := os.Args[1]
	iterations, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal(err)
	}

	m := utils.CreateMatrix(iterations)
	sum := 0
	for i := 0; i < iterations; i++ {
		for j := 0; j < iterations; j++ {
			sum += m[j][i]
		}
	}
	fmt.Printf("%v\n", sum)
}
