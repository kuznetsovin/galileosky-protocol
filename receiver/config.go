package main

/*
Описание конфигурационного файла
*/

import (
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/labstack/gommon/log"
)

type settings struct {
	Host       string
	Port       string
	ConLiveSec int    `toml:"con_live_sec"`
	LogLevel   string `toml:"log_level"`
}

func (c *settings) Load(confPath string) error {
	if _, err := toml.DecodeFile(confPath, c); err != nil {
		return fmt.Errorf("Ошибка разбора файла настроек: %v", err)
	}

	return nil
}

func (c *settings) getListenAddress() string {
	return c.Host + ":" + c.Port
}

func (c *settings) getLogLevel() log.Lvl {
	var lvl log.Lvl

	switch c.LogLevel {
	case "DEBUG":
		lvl = log.DEBUG
		break
	case "INFO":
		lvl = log.INFO
		break
	case "WARN":
		lvl = log.WARN
		break
	case "ERROR":
		lvl = log.ERROR
		break
	default:
		lvl = log.INFO
	}
	return lvl
}

func (c *settings) getEmptyConnTTL() time.Duration {
	return time.Duration(c.ConLiveSec) * time.Second
}
