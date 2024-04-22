.PHONY: build clean test

GO = $(shell which go)
BIN = ./bin/

.PHONY: deps
deps:
	${GO} mod tidy

build_amd64:
	GOARCH=amd64 ${GO} build -o ${BIN}/amd64/ main.go

build_arm64:
	GOARCH=arm64 ${GO} build -o ${BIN}/arm64/ main.go

build_linux_amd64:
	GOOS=linux GOARCH=amd64 ${GO} build -o ${BIN}/linux/amd64/ main.go

build:
	${GO} build -o ${BIN} main.go

.PHONY: clean
clean:
	rm -rf bin/* || true

.PHONY: test
test:
	${GO} test ./...
