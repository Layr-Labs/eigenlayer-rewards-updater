.PHONY: build clean test

GO = $(shell which go)
BIN = ./bin/

binary_name = payment-updater

all: build

.PHONY: deps
deps:
	${GO} install github.com/vektra/mockery/v2@v2.42.3
	${GO} mod tidy

build_amd64:
	GOARCH=amd64 ${GO} build -o ${BIN}/amd64/${binary_name} main.go

build_arm64:
	GOARCH=arm64 ${GO} build -o ${BIN}/arm64/${binary_name} main.go

build_linux_amd64:
	GOOS=linux GOARCH=amd64 ${GO} build -o ${BIN}/linux/amd64/${binary_name} main.go

build:
	${GO} build -o ${BIN}${binary_name} main.go

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

.PHONY: ci-test
ci-test: deps test

.PHONY: docker
docker:
	docker build -t payments-updater:latest .
