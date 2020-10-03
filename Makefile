NAME := primap
SRCS := $(shell find . -type f -name '*.go')

.DEFAULT_GOAL := bin/$(NAME)

bin/$(NAME): $(SRCS)
	go build -o bin/$(NAME)

.PHONY: clean
clean:
	rm -rf bin/*

.PHONY: test
test:
	go test -count=1 $${TEST_ARGS} ./...

.PHONY: testrace
testrace:
	go test -count=1 $${TEST_ARGS} -race ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: fmtci
fmtci:
	! gofmt -d . | grep '^'

.PHONY: lint
lint:
	golint -set_exit_status ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: test_all
test_all: test testrace fmt lint vet

.PHONY: go2ts
go2ts:
	tscriptify -package=github.com/sue445/primap/server/db -target=frontend/app/components/ShopEntity.ts ShopEntity
	npm run prettier:write
