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

	s := utils.CreateSlice(iterations)
	sum := 0
	for i := 0; i < iterations; i++ {
		sum += s[i]
	}
	fmt.Println(sum)
}
