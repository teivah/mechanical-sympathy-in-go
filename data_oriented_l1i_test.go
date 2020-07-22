package tests

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"runtime"
	"testing"
)

var resultInterface interface{}

const itx = 10

func TestA(t *testing.T) {
	fmt.Printf("%v\n", op1())
	fmt.Printf("%v\n", op2())
}

func BenchmarkIns2(b *testing.B) {
	s := make([]func() interface{}, itx)
	for i := 0; i < itx/2; i++ {
		s[i] = op1
	}
	for i := itx / 2; i < itx; i++ {
		s[i] = op2
	}
	runtime.GC()
	b.ResetTimer()

	var r interface{}
	for i := 0; i < b.N; i++ {
		for j := 0; j < itx; j++ {
			r = s[j]()
		}
	}
	resultInterface = r
}

func BenchmarkIns1(b *testing.B) {
	s := make([]func() interface{}, itx)
	for i := 0; i < itx; i += 2 {
		s[i] = op1
		s[i+1] = op2
	}
	runtime.GC()
	b.ResetTimer()

	var r interface{}
	for i := 0; i < b.N; i++ {
		for j := 0; j < itx; j++ {
			r = s[j]()
		}
	}
	resultInterface = r
}

func op1() interface{} {
	key := []byte("5e8487e6")
	src := []byte("hello world123563332")

	block, err := des.NewCipher(key)
	if err != nil {
		return err
	}
	bs := block.BlockSize()
	src = ZeroPadding(src, bs)
	// src = PKCS5Padding(src, bs)
	if len(src)%bs != 0 {
		return errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return out
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func op2() interface{} {
	text := []byte("My Super Secret Code Stuff")
	key := []byte("passphrasewhichneedstobe32bytes!")
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}
	return gcm.Seal(nonce, nonce, text, nil)
}
