#!/bin/sh

GOOS=linux GOARCH=amd64 go build -o build/bk-linux-amd64 main.go
GOOS=linux GOARCH=arm go build -o build/bk-linux-arm main.go
GOOS=darwin GOARCH=amd64 go build -o build/bk-darwin-amd64 main.go
GOOS=darwin GOARCH=arm64 go build -o build/bk-darwin-arm64 main.go