package crand

import (
	"bytes"
)

const charSet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const charSetLen = 62

//String returns a random string using the charset a-z A-Z 0-9
func String(size int) (string, error) {
	buf := new(bytes.Buffer)
	var in uint
	var err error
	for i := 0; i < size; i++ {
		if in, err = Uint(charSetLen); err != nil {
			return "", err
		}
		buf.WriteByte(charSet[in])
	}
	return buf.String(), nil
}
