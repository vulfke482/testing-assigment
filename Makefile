test:
	GOCONFIG=config/local_test/ go test -timeout 10s -race -p 1 ./...

test-cover:
	GOCONFIG=config/local_test/ go test -cover -p 1 ./...

format:
	gofmt -s -w **/*.go

init:
	chmod -R +x .githooks/
	mkdir -p .git/hooks/
	find .git/hooks -type l -exec rm {} \;
	find .githooks -type f -exec ln -sf ../../{} .git/hooks/ \;

lint:
	golangci-lint run --exclude-use-default=false --skip-dirs=vendor --disable-all --enable=goimports --enable=gosimple --enable=typecheck --enable=unused --enable=golint --enable=deadcode --enable=structcheck --enable=varcheck --enable=errcheck --enable=ineffassign --enable=govet --enable=staticcheck --enable=gofmt --deadline=3m ./...

.PHONY: test test-cover format init lint