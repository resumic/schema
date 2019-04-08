FILE=./VERSION
RESUMIC_VERSION = `cat $(FILE)`
GITCOMMIT = `git describe --tags --always --dirty`
build: ./cmd/version.go
	go build -ldflags "-X github.com/resumic/schema/cmd.Version=$(RESUMIC_VERSION) -X github.com/resumic/schema/cmd.GitCommit=$(GITCOMMIT)" -o resumic
