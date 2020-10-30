# Galileosky protocol receiver

Base implementation galileo server [protocol](docs/galileo-protocol.pdf) without cryptography and compression.

## Install

```
git clone https://github.com/kuznetsovin/galileosky-receiver
cd galileosky-receiver/tools && ./build-receiver.sh
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
