#/bin/bash

go test github.com/kuznetsovin/galileosky-protocol/pkg/packet -cover
go test github.com/kuznetsovin/galileosky-protocol/cmd/receiver -cover