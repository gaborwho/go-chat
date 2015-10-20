.PHONY : format build

build :
	go build -o bin/server src/*.go

format :
	gofmt -w src/*.go

