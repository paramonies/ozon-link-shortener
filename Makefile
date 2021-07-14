# Go related commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test ./...

.PHONY: build
build: 
	${GOBUILD} -v -o apiserver ./cmd/apiserver

.PHONY: swag
swag:
	swag init -g ./cmd/apiserver/main.go 

.PHONY: test
test:
	$(GOTEST)

# Generates a coverage report
.PHONY: cover
cover:
	${GOCMD} test -coverprofile=coverage.out ./... && ${GOCMD} tool cover -html=coverage.out


.DEFAULT_GOAL := build