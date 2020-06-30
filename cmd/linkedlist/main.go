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

	node := utils.CreateLinkedList(iterations)
	sum := 0
	for i := 0; i < iterations; i++ {
		sum += node.Value
		node = node.Next
	}

	fmt.Printf("%v\n", sum)
}
