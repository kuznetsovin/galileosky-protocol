package galileo

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUintTag_Parse(t *testing.T) {
	b := []byte{0x82}
	r := UintTag{}

	if assert.NoError(t, r.Parse(b)) {
		assert.Equal(t, UintTag{0x82}, r)
	}

	r = UintTag{}
	b = []byte{0x32, 0x00}
	if assert.NoError(t, r.Parse(b)) {
		assert.Equal(t, UintTag{50}, r)
	}
}

func TestStringTag_Parse(t *testing.T) {
	b := []byte{0x38, 0x36, 0x32, 0x30, 0x35, 0x37, 0x30, 0x34, 0x37, 0x37, 0x34, 0x35, 0x35, 0x33, 0x31}
	r := StringTag{}

	if assert.NoError(t, r.Parse(b)) {
		assert.Equal(t, StringTag{"862057047745531"}, r)
	}
}

func TestTimeTag_Parse(t *testing.T) {
	b := []byte{0x4E, 0x83, 0xFF, 0x5C}
	r := TimeTag{}

	if assert.NoError(t, r.Parse(b)) {
		assert.Equal(t,
			TimeTag{time.Date(2019, time.June, 11, 10, 32, 46, 0, time.UTC)},
			r,
		)
	}
}

func TestCoordTag_Parse(t *testing.T) {
	b := []byte{0x07, 0xC0, 0x0E, 0x32, 0x03, 0xB8, 0xD7, 0x2D, 0x05}
	r := CoordTag{}

	if assert.NoError(t, r.Parse(b)) {
		assert.Equal(t,
			CoordTag{7, 0, 86.890424, 53.612224},
			r,
		)
	}
}

func TestSpeedTag_Parse(t *testing.T) {
	b := []byte{0x5C, 0x00, 0x48, 0x08}
	r := SpeedTag{}

	if assert.NoError(t, r.Parse(b)) {
		assert.Equal(t,
			SpeedTag{9.2, 212},
			r,
		)
	}
}

func TestIntTag_Parse(t *testing.T) {
	b := []byte{0x00, 0x00}
	r := IntTag{}

	if assert.NoError(t, r.Parse(b)) {
		assert.Equal(t,
			IntTag{0},
			r,
		)
	}
}

func TestBitsTag_Parse(t *testing.T) {
	b := []byte{0x01, 0x3a}
	r := BitsTag{}

	if assert.NoError(t, r.Parse(b)) {
		assert.Equal(t,
			BitsTag{"0011101000000001"},
			r,
		)
	}
}
