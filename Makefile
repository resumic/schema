# Build Info Variables
FILE = ./VERSION
RESUMIC_VERSION = `cat $(FILE)`
GITCOMMIT = `git describe --tags --always --dirty`
CODEPATH = `go list -m`
SUFFIX = "-dev"
export GO111MODULE = on

# Build Rules

pre-build:
	@go get github.com/gobuffalo/packr/v2/packr2
	@packr2

build: pre-build
	@go build -ldflags "-X $(CODEPATH)/cmd.version=$(RESUMIC_VERSION)$(SUFFIX) -X $(CODEPATH)/cmd.gitCommit=$(GITCOMMIT)" -o resumic
	@packr2 clean

release: pre-build
	@go build -ldflags "-s -w -X $(CODEPATH)/cmd.version=$(RESUMIC_VERSION) -X $(CODEPATH)/cmd.gitCommit=$(GITCOMMIT)" -o resumic
	@packr2 clean

nopackr: # Builds without installing packr. Useful if offline.
	@go build -ldflags "-s -w -X $(CODEPATH)/cmd.version=$(RESUMIC_VERSION) -X $(CODEPATH)/cmd.gitCommit=$(GITCOMMIT)" -o resumic


test: build
	go test -v ./...

.PHONY: clean
clean: ./resumic
	@rm resumic