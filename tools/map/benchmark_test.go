package main

import (
	"testing"
)

const N = 30e6

func BenchmarkPointer(b *testing.B) {
	m := make(map[int32]*int32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := int32(i)
		m[n] = &n
	}
}

func BenchmarkInt32(b *testing.B) {
	m := make(map[int32]int32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := int32(i)
		m[n] = n
	}
}

func BenchmarkShared(b *testing.B) {
	shards := make([]map[int32]int32, 100)
	b.ResetTimer()
	for i := range shards {
		shards[i] = make(map[int32]int32)
	}
	for i := 0; i < b.N; i++ {
		n := int32(i)
		shards[i%100][n] = n
	}
}

func BenchmarkSlice(b *testing.B) {
	b.ResetTimer()
	type t struct {
		p, q int32
	}
	var s []t
	for i := 0; i < b.N; i++ {
		n := int32(i)
		s = append(s, t{n, n})
	}
}
