package benchmark

import (
	"math/rand"
	"testing"
)

const a = 10 ^ 10

var x [a]int

func init() {
	for i := 0; i < len(x); i++ {
		x[i] = rand.Intn(a)
	}
}

func length() {
	for i := 0; i < len(x); i++ {
		_ = i
	}
}

func lengthOne() {
	length := len(x)
	for i := 0; i < length; i++ {
		_ = i
	}
}

func BenchmarkLen(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		length()
	}
}

func BenchmarkLenOne(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lengthOne()
	}
}
