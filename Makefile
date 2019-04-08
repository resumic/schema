FILE=./VERSION
RESUMIC_VERSION = `cat $(FILE)`
GITCOMMIT = `git describe --tags --always --dirty`
CODEPATH=`go list -m`
build: ./cmd/version.go
	go get
	go build -ldflags "-X $(CODEPATH)/cmd.Version=$(RESUMIC_VERSION) -X $(CODEPATH)/cmd.GitCommit=$(GITCOMMIT)" -o resumic
clean: ./resumic
	@rm resumic