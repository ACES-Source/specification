BUILD_DATE = `date +%FT%T%z`
BUILD_USER = $(USER)@`hostname`
VERSION = `git describe --tags`

# command to build and run on the local OS.
GO_BUILD = go build

# tools
BINARY_CONTRACT_CLI=tokenized

GO_DIST_DIR=dist/golang

GO_DIST = CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_BUILD) -a -tags netgo -ldflags "-w -X main.buildVersion=$(VERSION) -X main.buildDate=$(BUILD_DATE) -X main.buildUser=$(BUILD_USER)"

all: prepare tools run-generate format test

run-win: prepare-win tools run-generate format-win

protobuf:
	protoc --proto_path=dist/protobuf --go_opt=paths=source_relative --go_out=dist/golang/actions --js_out=library=actions,binary:dist/typescript/protobuf/actions --python_out=dist/python/ dist/protobuf/actions.proto
	protoc --proto_path=dist/protobuf --go_opt=paths=source_relative --go_out=dist/golang/instruments --js_out=library=instruments,binary:dist/typescript/protobuf/instruments --python_out=dist/python/ dist/protobuf/instruments.proto
	protoc --proto_path=dist/protobuf --go_opt=paths=source_relative --go_out=dist/golang/messages --js_out=library=messages,binary:dist/typescript/protobuf/messages --python_out=dist/python/ dist/protobuf/messages.proto

generate-code:
	go run internal/cmd/generate/main.go
	goimports -w $(GO_DIST_DIR)

run-generate: generate-code protobuf
	goimports -w $(GO_DIST_DIR)

dist-cli:
	$(GO_DIST) -o dist/$(BINARY_CONTRACT_CLI) cmd/$(BINARY_CONTRACT_CLI)/main.go

format: format-go

lint: lint-go

test: test-go

test-all: test-clean-cache test

test-clean-cache:
	go clean -testcache

format-go:
	goimports -w $(GO_DIST_DIR)

lint-go:
	golint $(GO_DIST_DIR)
	go vet ./...

# run the tests with coverage
test-go:
	cd $(GO_DIST_DIR) && go test ./...

tools:
	go get golang.org/x/lint/golint
	go get golang.org/x/tools/cmd/goimports

prepare:
	mkdir -p tmp

format-win:
	goimports -w dist\golang\protocol\

prepare-win:
	mkdir tmp | echo tmp exists
