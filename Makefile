GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=stash
BINARY_UNIX=$(BINARY_NAME)_unix

build:
	$(GOBUILD) -o $(BINARY_NAME)

test:
	$(GOTEST) -v ./...

exec-dev-save:
	$(GOBUILD) -o $(BINARY_NAME); ./$(BINARY_NAME) save
exec-dev-pop:
	$(GOBUILD) -o $(BINARY_NAME); ./$(BINARY_NAME) pop
exec-dev-show:
	$(GOBUILD) -o $(BINARY_NAME); ./$(BINARY_NAME) show
exec-dev-list:
	$(GOBUILD) -o $(BINARY_NAME); ./$(BINARY_NAME) list