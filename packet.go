package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

type galileoParsePacket struct {
	Client              uint32         `json:"client"`
	PacketID            uint32         `json:"packet_id"`
	NavigationTimestamp int64          `json:"navigation_unix_time"`
	ReceivedTimestamp   int64          `json:"received_unix_time"`
	Latitude            float64        `json:"latitude"`
	Longitude           float64        `json:"longitude"`
	Speed               uint16         `json:"speed"`
	Pdop                uint16         `json:"pdop"`
	Hdop                uint16         `json:"hdop"`
	Vdop                uint16         `json:"vdop"`
	Nsat                uint8          `json:"nsat"`
	Ns                  uint16         `json:"ns"`
	Course              uint8          `json:"course"`
	AnSensors           []anSensor     `json:"an_sensors"`
	LiquidSensors       []liquidSensor `json:"liquid_sensors"`
}

type liquidSensor struct {
	SensorNumber uint8  `json:"sensor_number"`
	ValueMm      uint32 `json:"value_mm"`
	ValueL       uint32 `json:"value_l"`
}

type anSensor struct {
	SensorNumber uint8  `json:"sensor_number"`
	Value        uint32 `json:"value"`
}

type GalileoPaket struct {
	Header byte   `json:"header"`
	Length uint16 `json:"length"`
	Tags   tags   `json:"tags"`
	Crc16  uint16 `json:"crc"`
}

func (g *GalileoPaket) Decode(pkg []byte) error {
	var (
		err error
	)

	paketBodyLen := len(pkg) - 2

	g.Crc16 = binary.LittleEndian.Uint16(pkg[paketBodyLen:])

	if crc16(pkg[:paketBodyLen]) != g.Crc16 {
		return fmt.Errorf("Crc пакета не совпадает")
	}

	buf := bytes.NewReader(pkg[:paketBodyLen])

	if g.Header, err = buf.ReadByte(); err != nil {
		return fmt.Errorf("Ошибка чтения залоговка пакета: %v", err)
	}

	lenBytes := make([]byte, 2)
	if _, err = buf.Read(lenBytes); err != nil {
		return fmt.Errorf("Ошибка чтения длины пакета: %v", err)
	}

	g.Length = binary.LittleEndian.Uint16(lenBytes)

	lenBits := strconv.FormatUint(uint64(g.Length), 2)
	if len(lenBits) < 1 {
		return fmt.Errorf("Не корректная длина пакета: %v", err)
	}

	if lenBits[:1] == "1" {
		// если есть не отправленные данные, зануляем старший бит
		g.Length = g.Length << 1 >> 1
	}

	for buf.Len() > 0 {
		t := tag{}
		if t.Tag, err = buf.ReadByte(); err != nil {
			return fmt.Errorf("Ошибка чтения тэга: %v", err)
		}

		if tagInfo, ok := tagsTable[t.Tag]; ok {
			tagVal := make([]byte, tagInfo.Len)
			if _, err := buf.Read(tagVal); err != nil {
				return fmt.Errorf("Не удалось считать значение тега %x: %v", t.Tag, err)
			}
			if err := t.SetValue(tagInfo.Type, tagVal); err != nil {
				return err
			}
			g.Tags = append(g.Tags, t)
		} else {
			return fmt.Errorf("Не найдена информаци по тегу %x", t.Tag)
		}

	}

	return err
}

func (g *GalileoPaket) encode() error {
	return nil
}
