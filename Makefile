.PHONY: deps lint

deps:
	which gometalinter || (go get -u -v github.com/alecthomas/gometalinter && gometalinter --install)

lint:
	gometalinter --exclude=vendor --disable-all --enable=vet --enable=vetshadow --enable=gofmt ./...
