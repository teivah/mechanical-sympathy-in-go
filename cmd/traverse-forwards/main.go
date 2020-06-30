package main

import (
	"fmt"

	"github.com/teivah/mechanical-sympathy-in-go/cmd/utils"
)

func main() {
	s := utils.CreateSlice(utils.Iteration)
	sum := 0
	for i := 0; i < utils.Iteration; i++ {
		sum += s[i]
	}
	fmt.Println(sum)
}
