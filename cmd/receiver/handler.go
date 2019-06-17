package main

import (
	"encoding/binary"
	"github.com/kuznetsovin/galileosky-protocol/pkg/galileo"
	"io"
	"net"
	"time"
)

const headerLen = 3

func handleRecvPkg(conn net.Conn, ttl time.Duration) {
	var (
		recvPacket []byte
	)
	defer conn.Close()

	//packet.GalileoPaket
	logger.Warnf("Установлено соединение с %s", conn.RemoteAddr())

	for {
	Received:
		connTimer := time.NewTimer(ttl)

		// считываем заголовок пакета
		headerBuf := make([]byte, headerLen)
		_, err := conn.Read(headerBuf)

		switch err {
		case nil:
			connTimer.Reset(ttl)

			// если пакет не егтс формата закрываем соединение
			if headerBuf[0] != 0x01 {
				logger.Warnf("Пакет не соответствует формату Galileo. Закрыто соедиение %s", conn.RemoteAddr())
				return
			}

			// вычисляем длину пакета, 2 байта после тега
			pkgLen := binary.LittleEndian.Uint16(headerBuf[1:])
			pkgLen <<= 1
			pkgLen >>= 1
			pkgLen += 2

			// получаем концовку пакета
			buf := make([]byte, pkgLen)
			if _, err := io.ReadFull(conn, buf); err != nil {
				logger.Errorf("Ошибка при получении тела пакета: %v", err)
				return
			}

			// формируем порлный пакет
			recvPacket = append(headerBuf, buf...)
		case io.EOF:
			<-connTimer.C
			_ = conn.Close()
			logger.Warnf("Соединение %s закрыто по таймауту", conn.RemoteAddr())
			return
		default:
			logger.Errorf("Ошибка при получении: %v", err)
			return
		}

		logger.Debugf("Принят пакет: %X", recvPacket)
		pkg := galileo.Packet{}
		if err := pkg.Decode(recvPacket); err != nil {
			logger.Warn("Ошибка расшифровки пакета")
			logger.Error(err)
			return
		}

		crc := make([]byte, 2)
		binary.LittleEndian.PutUint16(crc, pkg.Crc16)
		resp := append([]byte{0x02}, crc...)

		if _, err = conn.Write(resp); err != nil {
			logger.Errorf("Ошибка отправки пакета подтверждения: %v", err)
		}

	}
}
