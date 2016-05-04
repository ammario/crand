package crand

import (
	"crypto/rand"
	"time"
)

//Read returns a byte slice containing cryptographicly random data
func Read(byt []byte) {
	bytesRead := 0
	var n int
	for bytesRead < len(byt) {
		n, _ = rand.Read(byt[bytesRead:])
		bytesRead += n
		if bytesRead < len(byt) {
			time.Sleep(time.Millisecond * 100)
		}
	}
}
