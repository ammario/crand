package crand

import (
	"fmt"
	"testing"
)

//Sorry, didn't feel like writing complicated dispersion tests :(

func TestStrings(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", String(5))
	}
	fmt.Printf(" <- Do they look random ?\n")
}
