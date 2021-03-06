package crand

import (
	"crypto/rand"
	"encoding/binary"
	"math"
)

var (
	maxUint = ^uint(0)
	maxInt  = int(maxUint >> 1)
)

//Uint returns a random uint
func Uint(max uint) uint {
	buf := make([]byte, 8)
	rangeEnd := maxUint - (maxUint % max)
	for do := true; do; do = uint(binary.BigEndian.Uint64(buf)) > rangeEnd {
		if _, err := rand.Read(buf); err != nil {
			panic(err)
		}
	}
	return uint(binary.BigEndian.Uint64(buf)) % max
}

//Uint8 returns a random uint8
func Uint8(max uint8) uint8 {
	buf := make([]byte, 1)
	rangeEnd := math.MaxUint8 - (math.MaxUint8 % max)
	for do := true; do; do = buf[0] > rangeEnd {
		if _, err := rand.Read(buf); err != nil {
			panic(err)
		}
	}
	return buf[0] % max
}

//Uint16 returns a random uint16
func Uint16(max uint16) uint16 {
	buf := make([]byte, 2)
	rangeEnd := math.MaxUint16 - (math.MaxUint16 % max)
	for do := true; do; do = binary.BigEndian.Uint16(buf) > rangeEnd {
		if _, err := rand.Read(buf); err != nil {
			panic(err)
		}
	}
	return binary.BigEndian.Uint16(buf) % max
}

//Uint32 returns a random uint32
func Uint32(max uint32) uint32 {
	buf := make([]byte, 4)
	rangeEnd := math.MaxUint32 - (math.MaxUint32 % max)
	for do := true; do; do = binary.BigEndian.Uint32(buf) > rangeEnd {
		if _, err := rand.Read(buf); err != nil {
			panic(err)
		}
	}
	return binary.BigEndian.Uint32(buf) % max
}

//Uint64 returns a random uint64
func Uint64(max uint64) uint64 {
	buf := make([]byte, 8)
	rangeEnd := math.MaxUint64 - (math.MaxUint64 % max)
	for do := true; do; do = binary.BigEndian.Uint64(buf) > rangeEnd {
		if _, err := rand.Read(buf); err != nil {
			panic(err)
		}
	}
	return binary.BigEndian.Uint64(buf) % max
}
