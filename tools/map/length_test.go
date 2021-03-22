package main

import "testing"

func BenchmarkLen(b *testing.B) {
	m := make(map[int32]int32)
	b.ResetTimer()
	for i := 0; i < 10e6; i++ {
		n := int32(i)
		m[n] = n
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = len(m)
	}
}
