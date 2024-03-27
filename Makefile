.PHONY: build clean tests

build:
	go build ./cmd/locknock

clean:
	rm locknock

tests:
	go test ./...	
	go build ./...
