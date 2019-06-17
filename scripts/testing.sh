#/bin/bash

go test github.com/kuznetsovin/galileosky-protocol/pkg/galileo -cover
go test github.com/kuznetsovin/galileosky-protocol/cmd/receiver -cover