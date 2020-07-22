package main

import "time"

func main() {
	ch := make(chan struct{})
	a := 0
	b := 0

	go func() {
		b = 1
		<-ch
		if a == 0 {
			panic("a")
		}
	}()

	go func() {
		a = 1
		ch <- struct{}{}
		if b == 0 {
			panic("b")
		}
	}()

	time.Sleep(4000 * time.Millisecond)

}
