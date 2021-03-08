package main

import "testing"

func BenchmarkDirectCall(b *testing.B) {
	c := &Cat{Name: "xdhuxc"}
	for n := 0; n < b.N; n++ {
		c.Quack()
	}
}

func BenchmarkDynamicDispatch(b *testing.B) {
	c := Duck(&Cat{Name: "xdhuxc"})
	for n := 0; n < b.N; n++ {
		c.Quack()
	}
}
