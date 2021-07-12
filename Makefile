# Go related commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test ./...

.PHONY: build
build: 
	${GOBUILD} -v -o apiserver ./cmd/apiserver

.DEFAULT_GOAL := build