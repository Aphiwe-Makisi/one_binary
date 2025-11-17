.PHONY build test clean

BINARY=goscripts
LDFLAGS=-X main.version=dev

build:
	go build -ldflags "$(LDFLAGS)" -o $(BINARY) ./cmd/$(BINARY)

test:
	go test ./...

clean:
	rm -f $(BINARY)