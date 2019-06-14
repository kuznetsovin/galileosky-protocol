package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUintTag_Parse(t *testing.T) {
	b := []byte{0x82}
	r := new(uintTag)

	if assert.NoError(t, r.Parse(b)) {
		assert.Equal(t, uintTag(0x82), *r)
	}

	r = new(uintTag)
	b = []byte{0x32, 0x00}
	if assert.NoError(t, r.Parse(b)) {
		assert.Equal(t, uintTag(50), *r)
	}
}

func TestStringTag_Parse(t *testing.T) {
	b := []byte{0x38, 0x36, 0x32, 0x30, 0x35, 0x37, 0x30, 0x34, 0x37, 0x37, 0x34, 0x35, 0x35, 0x33, 0x31}
	r := new(stringTag)

	if assert.NoError(t, r.Parse(b)) {
		assert.Equal(t, stringTag("862057047745531"), *r)
	}
}
