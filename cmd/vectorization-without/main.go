package main

import (
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
	it := 4 * iterations

	x := utils.CreateSlice(it)
	y := x
	res := x
	for i := 0; i < it; i++ {
		res[i] = x[i] + y[i]
	}
}
