.PHONY : format build

build :
	go build -o bin/server src/*.go

format :
	gofmt -w src/*.go

install-hooks :
	ln -siT ../hooks .git/hooks

