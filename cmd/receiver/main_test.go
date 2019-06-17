package main

import (
	"bytes"
	"io/ioutil"
	"net"
	"testing"
	"time"

	"github.com/labstack/gommon/log"
)

func TestServer(t *testing.T) {
	logger = log.New("-")
	logger.SetOutput(ioutil.Discard)

	srv := "127.0.0.1:5020"
	message := []byte{0x01, 0x17, 0x80, 0x01, 0x82, 0x02, 0x10, 0x03, 0x38, 0x36, 0x32, 0x30, 0x35, 0x37, 0x30,
		0x34, 0x37, 0x37, 0x34, 0x35, 0x35, 0x33, 0x31, 0x04, 0x32, 0x00, 0xB5, 0x48}
	response := []byte{0x02, 0xB5, 0x48}
	// запускаем сервер
	go func() {
		runServer(srv, 5*time.Second)
	}()

	conn, err := net.Dial("tcp", srv)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	_ = conn.SetWriteDeadline(time.Now().Add(2 * time.Second))
	_, _ = conn.Write(message)

	buf := make([]byte, 3)
	_, _ = conn.Read(buf)

	if !bytes.Equal(buf, response) {
		t.Errorf("Ответне совпадает: %X != %X ", buf, response)
	}
}
