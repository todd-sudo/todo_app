.PHONY:

build:
	go build -o .bin/main cmd/main.go

run:
	go run cmd/main.go

all: build
	go run cmd/main.go