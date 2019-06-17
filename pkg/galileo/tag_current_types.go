package galileo

import (
	"encoding/binary"
	"fmt"
	"time"
)

type tagParser interface {
	Parse(val []byte) error
}

type uintTag struct {
	uint64
}

func (u *uintTag) Parse(val []byte) error {
	switch size := len(val); {
	case size == 1:
		u.uint64 = uint64(val[0])
	case size == 2:
		u.uint64 = uint64(binary.LittleEndian.Uint16(val))
	default:
		return fmt.Errorf("Входной массив больше 2 байт: %x", val)
	}

	return nil
}

type stringTag struct {
	string
}

func (s *stringTag) Parse(val []byte) error {
	s.string = string(val)

	return nil
}

type timeTag struct {
	time.Time
}

func (s *timeTag) Parse(val []byte) error {
	secs := int64(binary.LittleEndian.Uint32(val))
	s.Time = time.Unix(secs, 0).UTC()

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
