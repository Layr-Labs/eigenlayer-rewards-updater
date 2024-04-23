.PHONY: build clean test

GO = $(shell which go)
BIN = ./bin/

all: build

.PHONY: deps
deps:
	${GO} install github.com/vektra/mockery/v2@v2.42.3
	${GO} mod tidy

build_amd64: mocks
	GOARCH=amd64 ${GO} build -o ${BIN}/amd64/ main.go

build_arm64: mocks
	GOARCH=arm64 ${GO} build -o ${BIN}/arm64/ main.go

build_linux_amd64: mocks
	GOOS=linux GOARCH=amd64 ${GO} build -o ${BIN}/linux/amd64/ main.go

build: mocks
	${GO} build -o ${BIN} main.go

.PHONY: mocks
mocks:
	mockery --all --case snake

.PHONY: clean
clean:
	rm -rf bin/* || true
	rm -rf mocks/* || true

.PHONY: test
test:
	${GO} test ./...
