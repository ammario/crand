package crand

import (
	"fmt"
	"testing"
)

//Sorry, didn't feel like writing complicated dispersion tests :(

func TestStrings(t *testing.T) {
	var str string
	var err error
	for i := 0; i < 10; i++ {
		if str, err = String(5); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%v ", str)
		}
	}
	fmt.Printf(" <- Do they look random ?\n")
}
