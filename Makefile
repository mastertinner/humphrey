.PHONY: all lint test build-docker deploy-cf clean

.EXPORT_ALL_VARIABLES:
GO111MODULE = on

all:
	go build -o bin/humphrey ./cmd/humphrey
	./bin/humphrey

lint:
	golangci-lint run

test:
	go test -race -cover ./...

clean:
	rm -rf bin

