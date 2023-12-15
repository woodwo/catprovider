GO=go
TARGET_DIR?=$(PWD)/.build

.PHONY: build
build: 
	@GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build -o $(TARGET_DIR)/client ./cmd/*.go

.PHONY: run
run: build 
	@$(TARGET_DIR)/client

.PHONY: test
test:  
	@$(GO) test -v ./...

.PHONY: tidy
tidy:
	go mod tidy
%:
	@:
