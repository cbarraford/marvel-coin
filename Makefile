PWD=$(shell pwd)
export GOPATH=${HOME}

ifndef TARGET
	TARGET=$(shell go list ./... | grep -v /vendor/)
endif

.PHONY: get build start run test test-cli lint docs

get:
	go get -v ${TARGET}
	go get -u github.com/kardianos/govendor

BUILD_ARGS=-v
build-internal:
	go get -t ${BUILD_ARGS} ${TARGET}

start-internal:
	go run main.go

test-internal:
	go test ${TARGET}

test-cover-internal:
	./scripts/cover.sh

vet-internal:
	go vet ${TARGET}

fmt-check-internal:
	go fmt ${TARGET}

build:
	make build-internal

start:
	make start-internal

run: build start

test:
	make test-internal

test-cover:
	make test-cover-internal

lint-internal:
	make fmt-check-internal vet-internal

lint:
	make lint-internal
