build:
	go build -o lexiquery cmd/main.go

test:
	go test -v ./...

clean:
	rm -rf lexiquery

install:
	make build
	sudo cp lexiquery /usr/local/bin

.PHONY: build test clean