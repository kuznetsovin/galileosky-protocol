package galileo

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFisrPaket_Decode(t *testing.T) {
	testPaketBin := []byte{0x01, 0x17, 0x80, 0x01, 0x82, 0x02, 0x10, 0x03, 0x38, 0x36, 0x32, 0x30, 0x35, 0x37, 0x30,
		0x34, 0x37, 0x37, 0x34, 0x35, 0x35, 0x33, 0x31, 0x04, 0x32, 0x00, 0xB5, 0x48}

	testPaket := Packet{
		Header: 0x01,
		Length: 23,
		Tags: tags{
			tag{0x01, &UintTag{130}},
			tag{0x02, &UintTag{16}},
			tag{0x03, &StringTag{"862057047745531"}},
			tag{0x04, &UintTag{50}},
		},
		Crc16: 18613,
	}

	p := Packet{}
	if assert.NoError(t, p.Decode(testPaketBin)) {
		assert.Equal(t, testPaket, p)
	}
}

func TestPacket_Decode(t *testing.T) {
	testPaketBin := []byte{0x01, 0xE7, 0x83, 0x10, 0x00, 0x00, 0x20, 0x4E, 0x83, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x01,
		0x3A, 0x41, 0x3F, 0x30, 0x42, 0xFE, 0x0E, 0x10, 0x01, 0x00, 0x20, 0x8D, 0x83, 0xFF, 0x5C, 0x30, 0xF3, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40,
		0x00, 0x3A, 0x41, 0x3B, 0x2E, 0x42, 0x08, 0x0F, 0x10, 0x02, 0x00, 0x20, 0x06, 0x84, 0xFF, 0x5C, 0x30, 0xF3,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00,
		0x40, 0x00, 0x3A, 0x41, 0x2E, 0x30, 0x42, 0x22, 0x0F, 0x10, 0x03, 0x00, 0x20, 0x6B, 0x84, 0xFF, 0x5C, 0x30,
		0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35,
		0x00, 0x40, 0x00, 0x3A, 0x41, 0x45, 0x30, 0x42, 0x39, 0x0F, 0x10, 0x04, 0x00, 0x20, 0x76, 0x84, 0xFF, 0x5C,
		0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00,
		0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0xDF, 0x2D, 0x42, 0x3A, 0x0F, 0x10, 0x05, 0x00, 0x20, 0x81, 0x84, 0xFF,
		0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00,
		0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x5B, 0x30, 0x42, 0x3B, 0x0F, 0x10, 0x06, 0x00, 0x20, 0x8B, 0x84,
		0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34,
		0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x2B, 0x2E, 0x42, 0x3E, 0x0F, 0x10, 0x07, 0x00, 0x20, 0x96,
		0x84, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00,
		0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0xF2, 0x2D, 0x42, 0x3F, 0x0F, 0x10, 0x08, 0x00, 0x20,
		0xA1, 0x84, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00,
		0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x83, 0x2E, 0x42, 0x3E, 0x0F, 0x10, 0x09, 0x00,
		0x20, 0xAC, 0x84, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00,
		0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x49, 0x30, 0x42, 0x42, 0x0F, 0x10, 0x0A,
		0x00, 0x20, 0xB7, 0x84, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00,
		0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x3A, 0x30, 0x42, 0x43, 0x0F, 0x10,
		0x0B, 0x00, 0x20, 0xC2, 0x84, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33,
		0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x3D, 0x30, 0x42, 0x43, 0x0F,
		0x10, 0x0C, 0x00, 0x20, 0xCD, 0x84, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x1F, 0x30, 0x42, 0x45,
		0x0F, 0x10, 0x0D, 0x00, 0x20, 0xD8, 0x84, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x2A, 0x2D, 0x42,
		0x46, 0x0F, 0x10, 0x0E, 0x00, 0x20, 0xE3, 0x84, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x33, 0x30,
		0x42, 0x48, 0x0F, 0x10, 0x0F, 0x00, 0x20, 0xEE, 0x84, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x47,
		0x30, 0x42, 0x48, 0x0F, 0x10, 0x10, 0x00, 0x20, 0xF9, 0x84, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41,
		0x25, 0x30, 0x42, 0x47, 0x0F, 0x10, 0x11, 0x00, 0x20, 0x04, 0x85, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A,
		0x41, 0x38, 0x30, 0x42, 0x4A, 0x0F, 0x10, 0x12, 0x00, 0x20, 0x0F, 0x85, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00,
		0x3A, 0x41, 0x0A, 0x2E, 0x42, 0x4B, 0x0F, 0x10, 0x13, 0x00, 0x20, 0x1A, 0x85, 0xFF, 0x5C, 0x30, 0xF3, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40,
		0x00, 0x3A, 0x41, 0x36, 0x30, 0x42, 0x4B, 0x0F, 0x10, 0x14, 0x00, 0x20, 0x25, 0x85, 0xFF, 0x5C, 0x30, 0xF3,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00,
		0x40, 0x00, 0x3A, 0x41, 0x3C, 0x30, 0x42, 0x4F, 0x0F, 0x10, 0x15, 0x00, 0x20, 0x30, 0x85, 0xFF, 0x5C, 0x30,
		0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35,
		0x00, 0x40, 0x00, 0x3A, 0x41, 0x42, 0x30, 0x42, 0x50, 0x0F, 0x10, 0x16, 0x00, 0x20, 0x3B, 0x85, 0xFF, 0x5C,
		0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00,
		0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x3C, 0x30, 0x42, 0x50, 0x0F, 0x10, 0x17, 0x00, 0x20, 0x46, 0x85, 0xFF,
		0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00,
		0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x42, 0x30, 0x42, 0x50, 0x0F, 0x10, 0x18, 0x00, 0x20, 0x51, 0x85,
		0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34,
		0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x4F, 0x30, 0x42, 0x50, 0x0F, 0x10, 0x19, 0x00, 0x20, 0x5C,
		0x85, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00,
		0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x3D, 0x30, 0x42, 0x51, 0x0F, 0x10, 0x1A, 0x00, 0x20,
		0x67, 0x85, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00,
		0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x2B, 0x30, 0x42, 0x52, 0x0F, 0xA5, 0x2D}

	p := Packet{}
	assert.NoError(t, p.Decode(testPaketBin))

}

func TestFisrPaket_Marshal(t *testing.T) {
	testPaketBin := []byte{0x01, 0x17, 0x80, 0x01, 0x82, 0x02, 0x10, 0x03, 0x38, 0x36, 0x32, 0x30, 0x35, 0x37, 0x30,
		0x34, 0x37, 0x37, 0x34, 0x35, 0x35, 0x33, 0x31, 0x04, 0x32, 0x00, 0xB5, 0x48}
	jsonStr := []byte(`{"header":1,"length":23,"tags":[{"tag":1,"value":{"val":130}},{"tag":2,"value":{"val":16}},{"tag":3,"value":{"val":"862057047745531"}},{"tag":4,"value":{"val":50}}],"crc":18613}`)

	p := Packet{}
	if assert.NoError(t, p.Decode(testPaketBin)) {
		r, err := json.Marshal(p)
		if assert.NoError(t, err) {
			assert.Equal(t, jsonStr, r)
		}
	}
}

func TestPaket_Marshal(t *testing.T) {
	testPaketBin := []byte{0x01, 0xE7, 0x83, 0x10, 0x00, 0x00, 0x20, 0x4E, 0x83, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x01,
		0x3A, 0x41, 0x3F, 0x30, 0x42, 0xFE, 0x0E, 0x10, 0x01, 0x00, 0x20, 0x8D, 0x83, 0xFF, 0x5C, 0x30, 0xF3, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40,
		0x00, 0x3A, 0x41, 0x3B, 0x2E, 0x42, 0x08, 0x0F, 0x10, 0x02, 0x00, 0x20, 0x06, 0x84, 0xFF, 0x5C, 0x30, 0xF3,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00,
		0x40, 0x00, 0x3A, 0x41, 0x2E, 0x30, 0x42, 0x22, 0x0F, 0x10, 0x03, 0x00, 0x20, 0x6B, 0x84, 0xFF, 0x5C, 0x30,
		0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35,
		0x00, 0x40, 0x00, 0x3A, 0x41, 0x45, 0x30, 0x42, 0x39, 0x0F, 0x10, 0x04, 0x00, 0x20, 0x76, 0x84, 0xFF, 0x5C,
		0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00,
		0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0xDF, 0x2D, 0x42, 0x3A, 0x0F, 0x10, 0x05, 0x00, 0x20, 0x81, 0x84, 0xFF,
		0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00,
		0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x5B, 0x30, 0x42, 0x3B, 0x0F, 0x10, 0x06, 0x00, 0x20, 0x8B, 0x84,
		0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34,
		0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x2B, 0x2E, 0x42, 0x3E, 0x0F, 0x10, 0x07, 0x00, 0x20, 0x96,
		0x84, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00,
		0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0xF2, 0x2D, 0x42, 0x3F, 0x0F, 0x10, 0x08, 0x00, 0x20,
		0xA1, 0x84, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00,
		0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x83, 0x2E, 0x42, 0x3E, 0x0F, 0x10, 0x09, 0x00,
		0x20, 0xAC, 0x84, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00,
		0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x49, 0x30, 0x42, 0x42, 0x0F, 0x10, 0x0A,
		0x00, 0x20, 0xB7, 0x84, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00,
		0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x3A, 0x30, 0x42, 0x43, 0x0F, 0x10,
		0x0B, 0x00, 0x20, 0xC2, 0x84, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33,
		0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x3D, 0x30, 0x42, 0x43, 0x0F,
		0x10, 0x0C, 0x00, 0x20, 0xCD, 0x84, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x1F, 0x30, 0x42, 0x45,
		0x0F, 0x10, 0x0D, 0x00, 0x20, 0xD8, 0x84, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x2A, 0x2D, 0x42,
		0x46, 0x0F, 0x10, 0x0E, 0x00, 0x20, 0xE3, 0x84, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x33, 0x30,
		0x42, 0x48, 0x0F, 0x10, 0x0F, 0x00, 0x20, 0xEE, 0x84, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x47,
		0x30, 0x42, 0x48, 0x0F, 0x10, 0x10, 0x00, 0x20, 0xF9, 0x84, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41,
		0x25, 0x30, 0x42, 0x47, 0x0F, 0x10, 0x11, 0x00, 0x20, 0x04, 0x85, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A,
		0x41, 0x38, 0x30, 0x42, 0x4A, 0x0F, 0x10, 0x12, 0x00, 0x20, 0x0F, 0x85, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00,
		0x3A, 0x41, 0x0A, 0x2E, 0x42, 0x4B, 0x0F, 0x10, 0x13, 0x00, 0x20, 0x1A, 0x85, 0xFF, 0x5C, 0x30, 0xF3, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40,
		0x00, 0x3A, 0x41, 0x36, 0x30, 0x42, 0x4B, 0x0F, 0x10, 0x14, 0x00, 0x20, 0x25, 0x85, 0xFF, 0x5C, 0x30, 0xF3,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35, 0x00,
		0x40, 0x00, 0x3A, 0x41, 0x3C, 0x30, 0x42, 0x4F, 0x0F, 0x10, 0x15, 0x00, 0x20, 0x30, 0x85, 0xFF, 0x5C, 0x30,
		0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x35,
		0x00, 0x40, 0x00, 0x3A, 0x41, 0x42, 0x30, 0x42, 0x50, 0x0F, 0x10, 0x16, 0x00, 0x20, 0x3B, 0x85, 0xFF, 0x5C,
		0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00,
		0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x3C, 0x30, 0x42, 0x50, 0x0F, 0x10, 0x17, 0x00, 0x20, 0x46, 0x85, 0xFF,
		0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00,
		0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x42, 0x30, 0x42, 0x50, 0x0F, 0x10, 0x18, 0x00, 0x20, 0x51, 0x85,
		0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00, 0x34,
		0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x4F, 0x30, 0x42, 0x50, 0x0F, 0x10, 0x19, 0x00, 0x20, 0x5C,
		0x85, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x00,
		0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x3D, 0x30, 0x42, 0x51, 0x0F, 0x10, 0x1A, 0x00, 0x20,
		0x67, 0x85, 0xFF, 0x5C, 0x30, 0xF3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00,
		0x00, 0x34, 0x00, 0x00, 0x35, 0x00, 0x40, 0x00, 0x3A, 0x41, 0x2B, 0x30, 0x42, 0x52, 0x0F, 0xA5, 0x2D}
	jsonStr := []byte(`{"header":1,"length":999,"tags":[{"tag":16,"value":{"val":0}},{"tag":32,"value":{"val":"2019-06-11T10:32:46Z"}},{"tag":48,"value":{"nsat":3,"is_valid":15,"longitude":0,"latitude":0}},{"tag":51,"value":{"speed":0,"course":0}},{"tag":52,"value":{"val":0}},{"tag":53,"value":{"val":0}},{"tag":64,"value":{"val":"0011101000000001"}},{"tag":65,"value":{"val":12351}},{"tag":66,"value":{"val":3838}},{"tag":16,"value":{"val":1}},{"tag":32,"value":{"val":"2019-06-11T10:33:49Z"}},{"tag":48,"value":{"nsat":3,"is_valid":15,"longitude":0,"latitude":0}},{"tag":51,"value":{"speed":0,"course":0}},{"tag":52,"value":{"val":0}},{"tag":53,"value":{"val":0}},{"tag":64,"value":{"val":"0011101000000000"}},{"tag":65,"value":{"val":11835}},{"tag":66,"value":{"val":3848}},{"tag":16,"value":{"val":2}},{"tag":32,"value":{"val":"2019-06-11T10:35:50Z"}},{"tag":48,"value":{"nsat":3,"is_valid":15,"longitude":0,"latitude":0}},{"tag":51,"value":{"speed":0,"course":0}},{"tag":52,"value":{"val":0}},{"tag":53,"value":{"val":0}},{"tag":64,"value":{"val":"0011101000000000"}},{"tag":65,"value":{"val":12334}},{"tag":66,"value":{"val":3874}},{"tag":16,"value":{"val":3}},{"tag":32,"value":{"val":"2019-06-11T10:37:31Z"}},{"tag":48,"value":{"nsat":3,"is_valid":15,"longitude":0,"latitude":0}},{"tag":51,"value":{"speed":0,"course":0}},{"tag":52,"value":{"val":0}},{"tag":53,"value":{"val":0}},{"tag":64,"value":{"val":"0011101000000000"}},{"tag":65,"value":{"val":12357}},{"tag":66,"value":{"val":3897}},{"tag":16,"value":{"val":4}},{"tag":32,"value":{"val":"2019-06-11T10:37:42Z"}},{"tag":48,"value":{"nsat":3,"is_valid":15,"longitude":0,"latitude":0}},{"tag":51,"value":{"speed":0,"course":0}},{"tag":52,"value":{"val":0}},{"tag":53,"value":{"val":0}},{"tag":64,"value":{"val":"0011101000000000"}},{"tag":65,"value":{"val":11743}},{"tag":66,"value":{"val":3898}},{"tag":16,"value":{"val":5}},{"tag":32,"value":{"val":"2019-06-11T10:37:53Z"}},{"tag":48,"value":{"nsat":3,"is_valid":15,"longitude":0,"latitude":0}},{"tag":51,"value":{"speed":0,"course":0}},{"tag":52,"value":{"val":0}},{"tag":53,"value":{"val":0}},{"tag":64,"value":{"val":"0011101000000000"}},{"tag":65,"value":{"val":12379}},{"tag":66,"value":{"val":3899}},{"tag":16,"value":{"val":6}},{"tag":32,"value":{"val":"2019-06-11T10:38:03Z"}},{"tag":48,"value":{"nsat":3,"is_valid":15,"longitude":0,"latitude":0}},{"tag":51,"value":{"speed":0,"course":0}},{"tag":52,"value":{"val":0}},{"tag":53,"value":{"val":0}},{"tag":64,"value":{"val":"0011101000000000"}},{"tag":65,"value":{"val":11819}},{"tag":66,"value":{"val":3902}},{"tag":16,"value":{"val":7}},{"tag":32,"value":{"val":"2019-06-11T10:38:14Z"}},{"tag":48,"value":{"nsat":3,"is_valid":15,"longitude":0,"latitude":0}},{"tag":51,"value":{"speed":0,"course":0}},{"tag":52,"value":{"val":0}},{"tag":53,"value":{"val":0}},{"tag":64,"value":{"val":"0011101000000000"}},{"tag":65,"value":{"val":11762}},{"tag":66,"value":{"val":3903}},{"tag":16,"value":{"val":8}},{"tag":32,"value":{"val":"2019-06-11T10:38:25Z"}},{"tag":48,"value":{"nsat":3,"is_valid":15,"longitude":0,"latitude":0}},{"tag":51,"value":{"speed":0,"course":0}},{"tag":52,"value":{"val":0}},{"tag":53,"value":{"val":0}},{"tag":64,"value":{"val":"0011101000000000"}},{"tag":65,"value":{"val":11907}},{"tag":66,"value":{"val":3902}},{"tag":16,"value":{"val":9}},{"tag":32,"value":{"val":"2019-06-11T10:38:36Z"}},{"tag":48,"value":{"nsat":3,"is_valid":15,"longitude":0,"latitude":0}},{"tag":51,"value":{"speed":0,"course":0}},{"tag":52,"value":{"val":0}},{"tag":53,"value":{"val":0}},{"tag":64,"value":{"val":"0011101000000000"}},{"tag":65,"value":{"val":12361}},{"tag":66,"value":{"val":3906}},{"tag":16,"value":{"val":10}},{"tag":32,"value":{"val":"2019-06-11T10:38:47Z"}},{"tag":48,"value":{"nsat":3,"is_valid":15,"longitude":0,"latitude":0}},{"tag":51,"value":{"speed":0,"course":0}},{"tag":52,"value":{"val":0}},{"tag":53,"value":{"val":0}},{"tag":64,"value":{"val":"0011101000000000"}},{"tag":65,"value":{"val":12346}},{"tag":66,"value":{"val":3907}},{"tag":16,"value":{"val":11}},{"tag":32,"value":{"val":"2019-06-11T10:38:58Z"}},{"tag":48,"value":{"nsat":3,"is_valid":15,"longitude":0,"latitude":0}},{"tag":51,"value":{"speed":0,"course":0}},{"tag":52,"value":{"val":0}},{"tag":53,"value":{"val":0}},{"tag":64,"value":{"val":"0011101000000000"}},{"tag":65,"value":{"val":12349}},{"tag":66,"value":{"val":3907}},{"tag":16,"value":{"val":12}},{"tag":32,"value":{"val":"2019-06-11T10:39:09Z"}},{"tag":48,"value":{"nsat":3,"is_valid":15,"longitude":0,"latitude":0}},{"tag":51,"value":{"speed":0,"course":0}},{"tag":52,"value":{"val":0}},{"tag":53,"value":{"val":0}},{"tag":64,"value":{"val":"0011101000000000"}},{"tag":65,"value":{"val":12319}},{"tag":66,"value":{"val":3909}},{"tag":16,"value":{"val":13}},{"tag":32,"value":{"val":"2019-06-11T10:39:20Z"}},{"tag":48,"value":{"nsat":3,"is_valid":15,"longitude":0,"latitude":0}},{"tag":51,"value":{"speed":0,"course":0}},{"tag":52,"value":{"val":0}},{"tag":53,"value":{"val":0}},{"tag":64,"value":{"val":"0011101000000000"}},{"tag":65,"value":{"val":11562}},{"tag":66,"value":{"val":3910}},{"tag":16,"value":{"val":14}},{"tag":32,"value":{"val":"2019-06-11T10:39:31Z"}},{"tag":48,"value":{"nsat":3,"is_valid":15,"longitude":0,"latitude":0}},{"tag":51,"value":{"speed":0,"course":0}},{"tag":52,"value":{"val":0}},{"tag":53,"value":{"val":0}},{"tag":64,"value":{"val":"0011101000000000"}},{"tag":65,"value":{"val":12339}},{"tag":66,"value":{"val":3912}},{"tag":16,"value":{"val":15}},{"tag":32,"value":{"val":"2019-06-11T10:39:42Z"}},{"tag":48,"value":{"nsat":3,"is_valid":15,"longitude":0,"latitude":0}},{"tag":51,"value":{"speed":0,"course":0}},{"tag":52,"value":{"val":0}},{"tag":53,"value":{"val":0}},{"tag":64,"value":{"val":"0011101000000000"}},{"tag":65,"value":{"val":12359}},{"tag":66,"value":{"val":3912}},{"tag":16,"value":{"val":16}},{"tag":32,"value":{"val":"2019-06-11T10:39:53Z"}},{"tag":48,"value":{"nsat":3,"is_valid":15,"longitude":0,"latitude":0}},{"tag":51,"value":{"speed":0,"course":0}},{"tag":52,"value":{"val":0}},{"tag":53,"value":{"val":0}},{"tag":64,"value":{"val":"0011101000000000"}},{"tag":65,"value":{"val":12325}},{"tag":66,"value":{"val":3911}},{"tag":16,"value":{"val":17}},{"tag":32,"value":{"val":"2019-06-11T10:40:04Z"}},{"tag":48,"value":{"nsat":3,"is_valid":15,"longitude":0,"latitude":0}},{"tag":51,"value":{"speed":0,"course":0}},{"tag":52,"value":{"val":0}},{"tag":53,"value":{"val":0}},{"tag":64,"value":{"val":"0011101000000000"}},{"tag":65,"value":{"val":12344}},{"tag":66,"value":{"val":3914}},{"tag":16,"value":{"val":18}},{"tag":32,"value":{"val":"2019-06-11T10:40:15Z"}},{"tag":48,"value":{"nsat":3,"is_valid":15,"longitude":0,"latitude":0}},{"tag":51,"value":{"speed":0,"course":0}},{"tag":52,"value":{"val":0}},{"tag":53,"value":{"val":0}},{"tag":64,"value":{"val":"0011101000000000"}},{"tag":65,"value":{"val":11786}},{"tag":66,"value":{"val":3915}},{"tag":16,"value":{"val":19}},{"tag":32,"value":{"val":"2019-06-11T10:40:26Z"}},{"tag":48,"value":{"nsat":3,"is_valid":15,"longitude":0,"latitude":0}},{"tag":51,"value":{"speed":0,"course":0}},{"tag":52,"value":{"val":0}},{"tag":53,"value":{"val":0}},{"tag":64,"value":{"val":"0011101000000000"}},{"tag":65,"value":{"val":12342}},{"tag":66,"value":{"val":3915}},{"tag":16,"value":{"val":20}},{"tag":32,"value":{"val":"2019-06-11T10:40:37Z"}},{"tag":48,"value":{"nsat":3,"is_valid":15,"longitude":0,"latitude":0}},{"tag":51,"value":{"speed":0,"course":0}},{"tag":52,"value":{"val":0}},{"tag":53,"value":{"val":0}},{"tag":64,"value":{"val":"0011101000000000"}},{"tag":65,"value":{"val":12348}},{"tag":66,"value":{"val":3919}},{"tag":16,"value":{"val":21}},{"tag":32,"value":{"val":"2019-06-11T10:40:48Z"}},{"tag":48,"value":{"nsat":3,"is_valid":15,"longitude":0,"latitude":0}},{"tag":51,"value":{"speed":0,"course":0}},{"tag":52,"value":{"val":0}},{"tag":53,"value":{"val":0}},{"tag":64,"value":{"val":"0011101000000000"}},{"tag":65,"value":{"val":12354}},{"tag":66,"value":{"val":3920}},{"tag":16,"value":{"val":22}},{"tag":32,"value":{"val":"2019-06-11T10:40:59Z"}},{"tag":48,"value":{"nsat":3,"is_valid":15,"longitude":0,"latitude":0}},{"tag":51,"value":{"speed":0,"course":0}},{"tag":52,"value":{"val":0}},{"tag":53,"value":{"val":0}},{"tag":64,"value":{"val":"0011101000000000"}},{"tag":65,"value":{"val":12348}},{"tag":66,"value":{"val":3920}},{"tag":16,"value":{"val":23}},{"tag":32,"value":{"val":"2019-06-11T10:41:10Z"}},{"tag":48,"value":{"nsat":3,"is_valid":15,"longitude":0,"latitude":0}},{"tag":51,"value":{"speed":0,"course":0}},{"tag":52,"value":{"val":0}},{"tag":53,"value":{"val":0}},{"tag":64,"value":{"val":"0011101000000000"}},{"tag":65,"value":{"val":12354}},{"tag":66,"value":{"val":3920}},{"tag":16,"value":{"val":24}},{"tag":32,"value":{"val":"2019-06-11T10:41:21Z"}},{"tag":48,"value":{"nsat":3,"is_valid":15,"longitude":0,"latitude":0}},{"tag":51,"value":{"speed":0,"course":0}},{"tag":52,"value":{"val":0}},{"tag":53,"value":{"val":0}},{"tag":64,"value":{"val":"0011101000000000"}},{"tag":65,"value":{"val":12367}},{"tag":66,"value":{"val":3920}},{"tag":16,"value":{"val":25}},{"tag":32,"value":{"val":"2019-06-11T10:41:32Z"}},{"tag":48,"value":{"nsat":3,"is_valid":15,"longitude":0,"latitude":0}},{"tag":51,"value":{"speed":0,"course":0}},{"tag":52,"value":{"val":0}},{"tag":53,"value":{"val":0}},{"tag":64,"value":{"val":"0011101000000000"}},{"tag":65,"value":{"val":12349}},{"tag":66,"value":{"val":3921}},{"tag":16,"value":{"val":26}},{"tag":32,"value":{"val":"2019-06-11T10:41:43Z"}},{"tag":48,"value":{"nsat":3,"is_valid":15,"longitude":0,"latitude":0}},{"tag":51,"value":{"speed":0,"course":0}},{"tag":52,"value":{"val":0}},{"tag":53,"value":{"val":0}},{"tag":64,"value":{"val":"0011101000000000"}},{"tag":65,"value":{"val":12331}},{"tag":66,"value":{"val":3922}}],"crc":11685}`)

	p := Packet{}
	if assert.NoError(t, p.Decode(testPaketBin)) {
		r, err := json.Marshal(p)
		if assert.NoError(t, err) {
			assert.Equal(t, string(jsonStr), string(r))
		}
	}
}
