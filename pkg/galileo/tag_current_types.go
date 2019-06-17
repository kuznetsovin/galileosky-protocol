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
	Val uint64 `json:"val"`
}

func (u *uintTag) Parse(val []byte) error {
	switch size := len(val); {
	case size == 1:
		u.Val = uint64(val[0])
	case size == 2:
		u.Val = uint64(binary.LittleEndian.Uint16(val))
	default:
		return fmt.Errorf("Входной массив больше 2 байт: %x", val)
	}

	return nil
}

type stringTag struct {
	Val string `json:"val"`
}

func (s *stringTag) Parse(val []byte) error {
	s.Val = string(val)

	return nil
}

type timeTag struct {
	Val time.Time `json:"val"`
}

func (t *timeTag) Parse(val []byte) error {
	secs := int64(binary.LittleEndian.Uint32(val))
	t.Val = time.Unix(secs, 0).UTC()

	return nil
}

type coordTag struct {
	Nsat      uint8
	isValid   uint8
	Longitude float64
	Latitude  float64
}

func (c *coordTag) Parse(val []byte) error {
	if len(val) != 9 {
		return fmt.Errorf(" Некорректная длин секции координат : %x", val)
	}

	flgByte := val[0]

	c.Latitude = float64(int32(binary.LittleEndian.Uint32(val[1:5]))) / float64(1000000)
	c.Longitude = float64(int32(binary.LittleEndian.Uint32(val[5:]))) / float64(1000000)

	c.Nsat = flgByte & 0xf
	c.isValid = flgByte >> 4

	return nil
}

type speedTag struct {
	Speed  float64
	Course uint16
}

func (s *speedTag) Parse(val []byte) error {
	if len(val) != 4 {
		return fmt.Errorf(" Некорректная длин секции скорости : %x", val)
	}

	s.Speed = float64(binary.LittleEndian.Uint16(val[:2])) / 10
	s.Course = binary.LittleEndian.Uint16(val[2:]) / 10
	return nil
}

type intTag struct {
	Val int `json:"val"`
}

func (u *intTag) Parse(val []byte) error {
	switch size := len(val); {
	case size == 1:
		u.Val = int(val[0])
	case size == 2:
		u.Val = int(binary.LittleEndian.Uint16(val))
	default:
		return fmt.Errorf("Входной массив больше 2 байт: %x", val)
	}

	return nil
}

type bitsTag struct {
	Val string `json:"val"`
}

func (b *bitsTag) Parse(val []byte) error {

	switch size := len(val); {
	case size == 1:
		b.Val = fmt.Sprintf("%08b", val[0])
	case size == 2:
		b.Val = fmt.Sprintf("%016b", binary.LittleEndian.Uint16(val))
	default:
		return fmt.Errorf("Входной массив больше 2 байт: %x", val)
	}

	return nil
}
