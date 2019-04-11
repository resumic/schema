# Build Info Variables
FILE=./VERSION
RESUMIC_VERSION = `cat $(FILE)`
GITCOMMIT = `git describe --tags --always --dirty`
CODEPATH=`go list -m`

# Build Rules

build: ./cmd/version.go
	@packr2 
	@go build -ldflags "-X $(CODEPATH)/cmd.version=$(RESUMIC_VERSION) -X $(CODEPATH)/cmd.gitCommit=$(GITCOMMIT)" -o resumic
	@packr2 clean

.PHONY: clean
clean: ./resumic
	@rm resumic