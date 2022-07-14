BINARY_NAME=iban-api
BINARY_PATH=./bin/$(BINARY_NAME)

WINDOWS=$(BINARY_PATH)-windows-amd64.exe
LINUX=$(BINARY_PATH)-linux-amd64
DARWIN=$(BINARY_PATH)-darwin-amd64

run:
	go run main.go
build:
	go build -o bin/main main.go
compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=amd64 go build -o ${WINDOWS} main.go
	GOOS=windows GOARCH=amd64 go build -o ${LINUX} main.go
	GOOS=darwin GOARCH=amd64 go build -o ${DARWIN} main.go

## OS env is only in windows
ifeq ($(OS),Windows_NT)
    os := Windows
else
    os := $(shell uname -s)
endif

## Build for all platforms specified
build:
ifeq ($(os), Windows)
	env GOOS=windows GOARCH=amd64 go build -o $(WINDOWS) main.go
endif
ifeq ($(os), Linux)
	env GOOS=linux GOARCH=amd64 go build -o $(LINUX) main.go
endif
ifeq ($(os), Darwin)
	env GOOS=darwin GOARCH=amd64 go build -o $(DARWIN) main.go
endif


run:
ifeq ($(os), Windows)
	$(WINDOWS)
endif
ifeq ($(os), Linux)
	$(LINUX)
endif
ifeq ($(os), Darwin)
	$(DARWIN)
endif

build_then_run: build run

test:
	cd ./pkg && go test -v