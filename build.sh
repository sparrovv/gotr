#!/bin/sh

for GOOS in linux windows darwin
do
  for GOARCH in amd64 386
  do
    GO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -o "bin/gotr_$GOOS-$GOARCH" gotr.go
  done
done
