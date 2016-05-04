package crand

import (
	"bytes"
)

const charSet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const charSetLen = 62

//String returns a random string using the charset a-z A-Z 0-9
func String(n int) string {
	buf := new(bytes.Buffer)
	for i := 0; i < n; i++ {
		buf.WriteByte(charSet[Uint(charSetLen)])
	}
	return buf.String()
}
