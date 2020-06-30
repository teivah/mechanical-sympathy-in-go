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
	i, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal(err)
	}

	s := utils.CreateSlice(i)
	sum := 0
	for i := i - 1; i >= 0; i-- {
		sum += s[i]
	}
	fmt.Println(sum)
}
