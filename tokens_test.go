package crand

import (
	"fmt"
	"testing"
)

//Sorry, didn't feel like writing complicated dispersion tests :(

func TestString(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", String(5))
	}
	fmt.Printf(" <- Do they look random ?\n")
}

func BenchmarkString20(b *testing.B) {
	var str string
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		str = String(20)
	}
	_ = str
}
