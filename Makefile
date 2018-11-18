.PHONY: coverage coverage-report install lint test
SHELL=/usr/bin/env bash -e -o pipefail

check-dependency = $(if $(shell command -v $(1)),,$(error Make sure $(1) is installed))

check-tools:
	@$(call check-dependency,gometalinter)

coverage:
	go test -v ./gcis/... -race -coverprofile=coverage.out -covermode=atomic

coverage-report:
	go tool cover -html=coverage.out

install:
	go get -u -v github.com/alecthomas/gometalinter && gometalinter --install

lint:check-tools
	gometalinter --exclude=vendor --disable-all --enable=vet --enable=vetshadow --enable=gofmt ./...

test:
	go test -v ./gcis/... -race
