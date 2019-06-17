package main

import (
	"net"
	"os"
	"time"

	"github.com/labstack/gommon/log"
)

var (
	config settings
	logger *log.Logger
)

func main() {
	logger = log.New("-")
	logger.SetHeader("${time_rfc3339_nano} ${short_file}:${line} ${level} -${message}")

	if len(os.Args) == 2 {
		if err := config.Load(os.Args[1]); err != nil {
			logger.Fatalf("Ошибка парсинга конфига: %v", err)
		}
	} else {
		logger.Fatalf("Не задан путь до конфига")
	}
	logger.SetLevel(config.getLogLevel())

	runServer(config.getListenAddress(), config.getEmptyConnTTL())
}

func runServer(srvAddress string, conTTL time.Duration) {
	l, err := net.Listen("tcp", srvAddress)
	if err != nil {
		logger.Fatalf("Не удалось открыть соединение: %v", err)
	}
	defer l.Close()

	logger.Infof("Запущен сервер %s...", srvAddress)
	for {
		conn, err := l.Accept()
		if err != nil {
			logger.Errorf("Ошибка соединения: %v", err)
		} else {
			go handleRecvPkg(conn, conTTL)
		}
	}
}
