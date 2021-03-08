package main

import (
	"strings"
	"testing"
)

var s = strings.Repeat("a", 1024)

func testStr2Bytes() {
	b := Str2Bytes(s)
	_ = Bytes2Str(b)
}

func testString() {
	b := []byte(s)
	_ = string(b)
}

func BenchmarkTestString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testString()
	}
}

func BenchmarkTestStr2Bytes(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testStr2Bytes()
	}
}
