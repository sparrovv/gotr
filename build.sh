#!/bin/sh

for GOOS in linux windows darwin
do
  for GOARCH in amd64 386
  do
    GO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -o "bin/gotr_$GOOS-$GOARCH" gotr.go
  done
done

# creating a package for homebrew
mkdir bin/bin
cp bin/gotr_darwin-amd64 bin/gotr
tar -czvf gotr-v0.2.0.tar.gz bin/gotr README.md
