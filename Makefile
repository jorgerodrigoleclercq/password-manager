.DEFAULT_GOAL := build

.PHONY: fmt vet build clean run debug

GOGC ?= 10
GODEBUG ?= gctrace=1

fmt: 
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -gcflags="-m"

run: build
	GOGC=$(GOGC) GODEBUG=$(GODEBUG) ./$(shell basename $(CURDIR))

debug: build
	GOGC=$(GOGC) GODEBUG=$(GODEBUG) ./$(shell basename $(CURDIR)) 2>&1 | tee gc_output.log

clean:
	go clean ./...
