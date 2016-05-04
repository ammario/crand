package crand

import (
	"fmt"
	"testing"
)

//Sorry, didn't feel like writing complicated dispersion tests :(

func TestInts(t *testing.T) {
	for i := 0; i < 10; i++ {
		if n, err := Uint8(10); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%v ", n)
		}
	}
	fmt.Printf(" <- Are they random and [0, 10) ?\n")
}
