# Build Info Variables
FILE=./VERSION
RESUMIC_VERSION = `cat $(FILE)`
GITCOMMIT = `git describe --tags --always --dirty`
CODEPATH=`go list -m`
DIRTY_SUFFIX="-dirty"

# Build Rules

build: 
	@packr2 
	@go build -o resumic
	@packr2 clean

dev: 
	@packr2 
	@go build -ldflags "-X $(CODEPATH)/cmd.version=$(RESUMIC_VERSION)$(DIRTY_SUFFIX) -X $(CODEPATH)/cmd.gitCommit=$(GITCOMMIT)" -o resumic
	@packr2 clean

release:
	@packr2
	@go build -ldflags "-X $(CODEPATH)/cmd.version=$(RESUMIC_VERSION) -X $(CODEPATH)/cmd.gitCommit=$(GITCOMMIT)" -o resumic
	@packr2 clean

.PHONY: clean
clean: ./resumic
	@rm resumic