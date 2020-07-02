package flate

import (
	"io/ioutil"
	"log"
	"testing"
)

func BenchmarkName(b *testing.B) {
	buf, err := ioutil.ReadFile("/Users/teivah/idea/mechanical-sympathy-in-go/cmd/flate/testdata/Isaac.Newton-Opticks.txt")
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
	w, err := NewWriter(ioutil.Discard, HuffmanOnly)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		w.Reset(ioutil.Discard)
		w.Write(buf1)
		w.Close()
	}
}
