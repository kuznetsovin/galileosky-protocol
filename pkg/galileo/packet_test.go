package galileo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGalileoPaket_Decode(t *testing.T) {
	testPaketBin := []byte{0x01, 0x17, 0x80, 0x01, 0x82, 0x02, 0x10, 0x03, 0x38, 0x36, 0x32, 0x30, 0x35, 0x37, 0x30,
		0x34, 0x37, 0x37, 0x34, 0x35, 0x35, 0x33, 0x31, 0x04, 0x32, 0x00, 0xB5, 0x48}

	v1 := uintTag(130)
	v2 := uintTag(16)
	v3 := stringTag("862057047745531")
	v4 := uintTag(50)

	testPaket := Packet{
		Header: 0x01,
		Length: 23,
		Tags: tags{
			tag{0x01, &v1},
			tag{0x02, &v2},
			tag{0x03, &v3},
			tag{0x04, &v4},
		},
		Crc16: 18613,
	}

	p := Packet{}
	if assert.NoError(t, p.Decode(testPaketBin)) {
		assert.Equal(t, testPaket, p)
	}
}
