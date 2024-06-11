#!/bin/bash
BIN=httptool
TARGET=cmd/cli.go
TIME=$(date -u "+%Y-%m-%d %H:%M:%S")

GOOS=linux GOARCH=amd64 go build -v -ldflags '-extldflags "-static"' -o bin/linux/${BIN} ${TARGET} && echo "${TIME}: build httptool for linux 64bits on bin/linux/${BIN}"
GOOS=windows GOARCH=amd64 go build -v -ldflags '-extldflags "-static"' -o bin/windows/${BIN}.exe ${TARGET} && echo "${TIME}: build httptool for windows 64bits on bin/windows/${BIN}.exe"

echo "Run ${BIN}:"
./bin/linux/${BIN}
