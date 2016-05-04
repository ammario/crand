package crand

import (
	"crypto/rand"
	"encoding/binary"
)

const (
	maxUint8  = 255
	maxUint16 = 65535
	maxUint32 = 4294967295
	maxUint64 = 1.8446744e19

	maxInt8  = 127
	maxInt16 = 32767
	maxInt32 = 2147483647
	maxInt64 = 9.223372e18
)

var (
	maxUint = ^uint(0)

	maxInt = int(maxUint >> 1)
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
	rangeEnd := maxUint8 - (maxUint8 % max)
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
	rangeEnd := maxUint16 - (maxUint16 % max)
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
	rangeEnd := maxUint32 - (maxUint32 % max)
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
	rangeEnd := maxUint64 - (maxUint64 % max)
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
	rangeEnd := maxInt8 - (maxInt8 % max)
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
	rangeEnd := maxInt16 - (maxInt16 % max)
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
	rangeEnd := maxInt32 - (maxInt32 % max)
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
	rangeEnd := maxInt64 - (maxInt64 % max)
	for do := true; do; do = int64(binary.BigEndian.Uint64(buf)) > rangeEnd {
		if _, err := rand.Read(buf); err != nil {
			return 0, err
		}
	}
	return int64(binary.BigEndian.Uint64(buf)) % max, nil
}
