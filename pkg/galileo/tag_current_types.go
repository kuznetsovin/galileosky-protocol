package galileo

import (
	"encoding/binary"
	"fmt"
)

type tagParser interface {
	Parse(val []byte) error
}

type uintTag uint64

func (u *uintTag) Parse(val []byte) error {
	var v uintTag

	switch size := len(val); {
	case size == 1:
		v = uintTag(val[0])
	case size == 2:
		v = uintTag(binary.LittleEndian.Uint16(val))
	default:
		return fmt.Errorf("Входной массив больше 2 байт: %x", val)
	}

	*u = v
	return nil
}

type stringTag string

func (s *stringTag) Parse(val []byte) error {
	*s = stringTag(string(val))

	return nil
}

type coordTag struct {
	Nsat      uint8
	isValid   uint8
	Longitude float64
	Latitude  float64
}

type speedTag struct {
	Speed  uint16
	Course uint16
}
