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
	y := utils.CreateSlice(it)
	res := utils.CreateSlice(it)
	for i := 0; i < it; i += 4 {
		res[i] = x[i] + y[i]
		res[i+1] = x[i+1] + y[i+1]
		res[i+2] = x[i+2] + y[i+2]
		res[i+3] = x[i+3] + y[i+3]
	}
}
