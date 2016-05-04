package crand

import (
	"fmt"
	"testing"
)

//Sorry, didn't feel like writing complicated dispersion tests :(

func TestInts(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", Uint8(10))
	}
	fmt.Printf(" <- Are they random and [0, 10) ?\n")
}
