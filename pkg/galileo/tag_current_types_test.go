package galileo

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUintTag_Parse(t *testing.T) {
	b := []byte{0x82}
	r := uintTag{}

	if assert.NoError(t, r.Parse(b)) {
		assert.Equal(t, uintTag{0x82}, r)
	}

	r = uintTag{}
	b = []byte{0x32, 0x00}
	if assert.NoError(t, r.Parse(b)) {
		assert.Equal(t, uintTag{50}, r)
	}
}

func TestStringTag_Parse(t *testing.T) {
	b := []byte{0x38, 0x36, 0x32, 0x30, 0x35, 0x37, 0x30, 0x34, 0x37, 0x37, 0x34, 0x35, 0x35, 0x33, 0x31}
	r := stringTag{}

	if assert.NoError(t, r.Parse(b)) {
		assert.Equal(t, stringTag{"862057047745531"}, r)
	}
}

func TestTimeTag_Parse(t *testing.T) {
	b := []byte{0x4E, 0x83, 0xFF, 0x5C}
	r := timeTag{}

	if assert.NoError(t, r.Parse(b)) {
		assert.Equal(t,
			timeTag{time.Date(2019, time.June, 11, 10, 32, 46, 0, time.UTC)},
			r,
		)
	}
}

func TestCoordTag_Parse(t *testing.T) {
	b := []byte{0x07, 0xC0, 0x0E, 0x32, 0x03, 0xB8, 0xD7, 0x2D, 0x05}
	r := coordTag{}

	if assert.NoError(t, r.Parse(b)) {
		assert.Equal(t,
			coordTag{7, 0, 86.890424, 53.612224},
			r,
		)
	}
}

func TestSpeedTag_Parse(t *testing.T) {
	b := []byte{0x5C, 0x00, 0x48, 0x08}
	r := speedTag{}

	if assert.NoError(t, r.Parse(b)) {
		assert.Equal(t,
			speedTag{9.2, 212},
			r,
		)
	}
}
