package galileo

import (
	"encoding/binary"
	"fmt"
	"time"
)

type tagParser interface {
	Parse(val []byte) error
}

//UintTag тип тэга беззнаковое целое
type UintTag struct {
	Val uint64 `json:"val"`
}

//Parse заполняет значение тэга
func (u *UintTag) Parse(val []byte) error {
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

//StringTag тип тэга строка
type StringTag struct {
	Val string `json:"val"`
}

//Parse заполняет значение тэга
func (s *StringTag) Parse(val []byte) error {
	s.Val = string(val)

	return nil
}

//TimeTag тип тэга время
type TimeTag struct {
	Val time.Time `json:"val"`
}

//Parse заполняет значение тэга
func (t *TimeTag) Parse(val []byte) error {
	secs := int64(binary.LittleEndian.Uint32(val))
	t.Val = time.Unix(secs, 0).UTC()

	return nil
}

//TimeTag тип тэга с координатами
type CoordTag struct {
	Nsat      uint8   `json:"nsat"`
	IsValid   uint8   `json:"is_valid"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

//Parse заполняет значение тэга
func (c *CoordTag) Parse(val []byte) error {
	if len(val) != 9 {
		return fmt.Errorf(" Некорректная длин секции координат : %x", val)
	}

	flgByte := val[0]

	c.Latitude = float64(int32(binary.LittleEndian.Uint32(val[1:5]))) / float64(1000000)
	c.Longitude = float64(int32(binary.LittleEndian.Uint32(val[5:]))) / float64(1000000)

	c.Nsat = flgByte & 0xf
	c.IsValid = flgByte >> 4

	return nil
}

//SpeedTag тип тэга со скоростью
type SpeedTag struct {
	Speed  float64 `json:"speed"`
	Course uint16  `json:"course"`
}

//Parse заполняет значение тэга
func (s *SpeedTag) Parse(val []byte) error {
	if len(val) != 4 {
		return fmt.Errorf(" Некорректная длин секции скорости : %x", val)
	}

	s.Speed = float64(binary.LittleEndian.Uint16(val[:2])) / 10
	s.Course = binary.LittleEndian.Uint16(val[2:]) / 10
	return nil
}

//IntTag тип тэга знаковго целого
type IntTag struct {
	Val int `json:"val"`
}

//Parse заполняет значение тэга
func (u *IntTag) Parse(val []byte) error {
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

//BitsTag тип тэга с битами
type BitsTag struct {
	Val string `json:"val"`
}

//Parse заполняет значение тэга
func (b *BitsTag) Parse(val []byte) error {

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
