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
func Uint(max uint) (uint, error) {
	buf := make([]byte, 8)
	rangeEnd := maxUint - (maxUint % max)
	for do := true; do; do = uint(binary.BigEndian.Uint64(buf)) > rangeEnd {
		if _, err := rand.Read(buf); err != nil {
			return 0, err
		}
	}
	return uint(binary.BigEndian.Uint64(buf)) % max, nil
}

//Uint8 returns a random uint8
func Uint8(max uint8) (uint8, error) {
	buf := make([]byte, 1)
	rangeEnd := math.MaxUint8 - (math.MaxUint8 % max)
	for do := true; do; do = buf[0] > rangeEnd {
		if _, err := rand.Read(buf); err != nil {
			return 0, err
		}
	}
	return buf[0] % max, nil
}

//Uint16 returns a random uint16
func Uint16(max uint16) (uint16, error) {
	buf := make([]byte, 2)
	rangeEnd := math.MaxUint16 - (math.MaxUint16 % max)
	for do := true; do; do = binary.BigEndian.Uint16(buf) > rangeEnd {
		if _, err := rand.Read(buf); err != nil {
			return 0, err
		}
	}
	return binary.BigEndian.Uint16(buf) % max, nil
}

//Uint32 returns a random uint32
func Uint32(max uint32) (uint32, error) {
	buf := make([]byte, 4)
	rangeEnd := math.MaxUint32 - (math.MaxUint32 % max)
	for do := true; do; do = binary.BigEndian.Uint32(buf) > rangeEnd {
		if _, err := rand.Read(buf); err != nil {
			return 0, err
		}
	}
	return binary.BigEndian.Uint32(buf) % max, nil
}

//Uint64 returns a random uint64
func Uint64(max uint64) (uint64, error) {
	buf := make([]byte, 8)
	rangeEnd := math.MaxUint64 - (math.MaxUint64 % max)
	for do := true; do; do = binary.BigEndian.Uint64(buf) > rangeEnd {
		if _, err := rand.Read(buf); err != nil {
			return 0, err
		}
	}
	return binary.BigEndian.Uint64(buf) % max, nil
}

//Int returns a random int
func Int(max int) (int, error) {
	buf := make([]byte, 8)
	rangeEnd := maxInt - (maxInt % max)
	for do := true; do; do = int(binary.BigEndian.Uint64(buf)) > rangeEnd {
		if _, err := rand.Read(buf); err != nil {
			return 0, err
		}
	}
	return int(binary.BigEndian.Uint64(buf)) % max, nil
}

//Int8 returns a random int8
func Int8(max int8) (int8, error) {
	buf := make([]byte, 1)
	rangeEnd := math.MaxInt8 - (math.MaxInt8 % max)
	for do := true; do; do = int8(buf[0]) > rangeEnd {
		if _, err := rand.Read(buf); err != nil {
			return 0, err
		}
	}
	return int8(buf[0]) % max, nil
}

//Int16 returns a random int16
func Int16(max int16) (int16, error) {
	buf := make([]byte, 2)
	rangeEnd := math.MaxInt16 - (math.MaxInt16 % max)
	for do := true; do; do = int16(binary.BigEndian.Uint16(buf)) > rangeEnd {
		if _, err := rand.Read(buf); err != nil {
			return 0, err
		}
	}
	return int16(binary.BigEndian.Uint16(buf)) % max, nil
}

//Int32 returns a random int32
func Int32(max int32) (int32, error) {
	buf := make([]byte, 4)
	rangeEnd := math.MaxInt32 - (math.MaxInt32 % max)
	for do := true; do; do = int32(binary.BigEndian.Uint32(buf)) > rangeEnd {
		if _, err := rand.Read(buf); err != nil {
			return 0, err
		}
	}
	return int32(binary.BigEndian.Uint32(buf)) % max, nil
}

//Int64 returns a random int64
func Int64(max int64) (int64, error) {
	buf := make([]byte, 4)
	rangeEnd := math.MaxInt64 - (math.MaxInt64 % max)
	for do := true; do; do = int64(binary.BigEndian.Uint64(buf)) > rangeEnd {
		if _, err := rand.Read(buf); err != nil {
			return 0, err
		}
	}
	return int64(binary.BigEndian.Uint64(buf)) % max, nil
}
