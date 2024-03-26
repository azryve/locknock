.PHONY: build clean tests

build:
	go build locknock/cmd/locknock

clean:
	rm locknock

tests:
	go test locknock/internal/locknock
