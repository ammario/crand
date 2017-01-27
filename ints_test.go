package crand

import (
	"fmt"
	"testing"
)

func TestInts(t *testing.T) {
	//ik it's shitty
	for i := 0; i < 10; i++ {
		n := Uint8(10)
		fmt.Printf("%v ", n)
	}
	fmt.Printf(" <- Are they random and [0, 10) ?\n")
}

func BenchmarkUint(b *testing.B) {
	var u uint8
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		u = Uint8(10)
	}
	_ = u
}
