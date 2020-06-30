package main

import (
	"fmt"

	"github.com/teivah/mechanical-sympathy-in-go/cmd/utils"
)

func main() {
	s := utils.CreateSlice(utils.Iteration)
	sum := 0
	for i := utils.Iteration - 1; i >= 0; i-- {
		sum += s[i]
	}
	fmt.Println(sum)
}
