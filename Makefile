.PHONY: build
build:
	go build -o bin/humphrey ./cmd/humphrey

.PHONY: run
run:
	go run cmd/humphrey/main.go

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -race -cover ./...

.PHONY: clean
clean:
	rm -rf bin
