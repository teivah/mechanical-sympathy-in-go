package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/teivah/mechanical-sympathy-in-go/cmd/flate"
)

func main() {
	iterations, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	n := int(1e4)
	buf1 := make([]byte, n)
	for i := 0; i < n; i += len(buf) {
		if len(buf) > n-i {
			buf = buf[:n-i]
		}
		copy(buf1[i:], buf)
	}
	buf = nil
	w, err := flate.NewWriter(ioutil.Discard, flate.HuffmanOnly)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < iterations; i++ {
		w.Reset(ioutil.Discard)
		w.Write(buf1)
		w.Close()
	}
}
