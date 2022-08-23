package test

import (
	"testing"
)

const (
	blocks    = 64
	blockSize = 1024
)

var block = make([]byte, blockSize)

func BenchmarkFori(b *testing.B) {
	a := make([]byte, blocks*blockSize)
	for n := 0; n < b.N; n++ {
		for i := 0; i < blocks; i++ {
			for j := 0; j < blockSize; j++ {
				a[i*blockSize+j] = block[j]
			}
		}
	}
}

func BenchmarkAppend(b *testing.B) {
	a := make([]byte, 0, blocks*blockSize)
	for n := 0; n < b.N; n++ {
		a = a[:0]
		for i := 0; i < blocks; i++ {
			a = append(a, block...)
		}
	}
}

func BenchmarkCopy(b *testing.B) {
	a := make([]byte, blocks*blockSize)
	for n := 0; n < b.N; n++ {
		for i := 0; i < blocks; i++ {
			copy(a[i*blockSize:], block)
		}
	}
}
