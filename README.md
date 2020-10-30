# Galileosky protocol receiver

Base implementation galileo server [protocol](docs/galileo-protocol.pdf) without cryptography and compression.

## Install

```
git clone https://github.com/kuznetsovin/galileosky-protocol
cd galileosky-protocol/tools && ./build-receiver.sh
```

## Run

```
./receiver config.toml
config.toml - configure file
```

## Config format

[srv]
host = "127.0.0.1"
port = "6000"
con_live_sec = 10
log_level = "DEBUG"

Parameters description:

```host``` - bind address
```port``` - bind port
```con_live_sec``` - if server not received data longer time in the parameter, then the connection is closed.
```log_level``` - logging level

## Using galileosky parsing library

You can use only parsing library without the server. For example:

```
package main 

import (
    "log"
    "github.com/kuznetsovin/galileosky-protocol/receiver/galileo"
)

func main() {
    pkg := []byte{0x01, 0x17, 0x80, 0x01, 0x82, 0x02, 0x10, 0x03, 0x38, 0x36, 0x32, 0x30, 0x35, 0x37, 0x30,
           		0x34, 0x37, 0x37, 0x34, 0x35, 0x35, 0x33, 0x31, 0x04, 0x32, 0x00, 0xB5, 0x48}
    result := galileo.Package{}

    if err := result.Decode(pkg); err != nil {
 		log.Fatal(err)
 	}
    
    log.Println("Package: ", result)
}
```