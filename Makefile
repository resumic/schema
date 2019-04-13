# Build Info Variables
FILE=./VERSION
RESUMIC_VERSION = `cat $(FILE)`
GITCOMMIT = `git describe --tags --always --dirty`
CODEPATH=`go list -m`
DIRTY_SUFFIX="-dirty"
export GO111MODULE=on

# Build Rules

build: 
	@go get github.com/gobuffalo/packr/v2/packr2
	@packr2 
	@go build -o resumic
	@packr2 clean

dev:
	@go get github.com/gobuffalo/packr/v2/packr2
	@packr2 
	@go build -ldflags "-X $(CODEPATH)/cmd.version=$(RESUMIC_VERSION)$(DIRTY_SUFFIX) -X $(CODEPATH)/cmd.gitCommit=$(GITCOMMIT)" -o resumic
	@packr2 clean

release:
	@go get github.com/gobuffalo/packr/v2/packr2	
	@packr2
	@go build -ldflags "-X $(CODEPATH)/cmd.version=$(RESUMIC_VERSION) -X $(CODEPATH)/cmd.gitCommit=$(GITCOMMIT)" -o resumic
	@packr2 clean

.PHONY: clean
clean: ./resumic
	@rm resumic