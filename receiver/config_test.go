package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestConfigLoad(t *testing.T) {
	cfg := `host = "127.0.0.1"
port = "5020"
con_live_sec = 10
log_level = "DEBUG"`

	file, err := ioutil.TempFile("/tmp", "config.toml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	if _, err = file.WriteString(cfg); err != nil {
		t.Fatal(err)
	}

	conf := settings{}
	if err = conf.Load(file.Name()); err != nil {
		t.Fatal(err)
	}

	testCfg := settings{
		Host:       "127.0.0.1",
		Port:       "5020",
		ConLiveSec: 10,
		LogLevel:   "DEBUG",
	}
	assert.Equal(t, testCfg, conf)
}
