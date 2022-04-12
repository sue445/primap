NAME := primap
SRCS := $(shell find . -type f -name '*.go')
TSCRIPTIFY_VERSION := $(shell grep 'github.com/tkrajina/typescriptify-golang-structs' go.mod | awk '{ print $$2 }')

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

.PHONY: tscriptify
tscriptify:
	go install github.com/tkrajina/typescriptify-golang-structs/tscriptify@$(TSCRIPTIFY_VERSION)
	tscriptify -package=github.com/sue445/primap/db -target=../frontend/app/components/ShopEntity.ts ShopEntity

.PHONY: go2ts
go2ts: tscriptify
	cd ../frontend/ && npm install && npm run prettier:write

.PHONY: go2ts_ci
go2ts_ci: tscriptify
	cd ../frontend/ && npm run prettier:write
