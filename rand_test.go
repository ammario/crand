package crand

import (
	"fmt"
	"testing"
)

//Sorry, didn't feel like writing complicated dispersion tests :(

func TestRand(t *testing.T) {
	byt := make([]byte, 5)
	for i := 0; i < 5; i++ {
		Read(byt)
		fmt.Printf("%x ", byt)
	}
	fmt.Println(" <- Do those look random?")
}
