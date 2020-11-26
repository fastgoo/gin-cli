BIN_NAME=gin-cli-server
BIN_PATH=./output

.PHONY: install
install:
	@echo "build $(BIN_NAME) >>>"
	rm -rf $(BIN_PATH)
	mkdir -p $(BIN_PATH)
	cp ./.env $(BIN_PATH)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o $(BIN_NAME) ./main.go
	#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_CMD) build -o $(API_BIN) ./main.go
	#upx --brute $(BIN_NAME)
	cp ./$(BIN_NAME) $(BIN_PATH)
	rm -rf ./$(BIN_NAME)
	@echo ">>> build $(BIN_NAME) complete"

.PHONY: all
all:
	@echo "build $(BIN_NAME) all platform >>>"
	rm -rf $(BIN_PATH)
	mkdir -p $(BIN_PATH)
	cp ./.env $(BIN_PATH)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags '-w -s' -o $(BIN_NAME)-mac ./main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o $(BIN_NAME)-linux ./main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags '-w -s' -o $(BIN_NAME)-windows ./main.go
	#upx --brute $(BIN_NAME)
	cp ./$(BIN_NAME)-* $(BIN_PATH)
	rm -rf ./$(BIN_NAME)-*
	@echo ">>> build $(BIN_NAME) all platform complete"

.PHONY: mac
mac:
	@echo "build $(BIN_NAME) mac platform >>>"
	rm -rf $(BIN_PATH)
	mkdir -p $(BIN_PATH)
	cp ./.env $(BIN_PATH)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags '-w -s' -o $(BIN_NAME) ./main.go
	#upx --brute $(BIN_NAME)
	cp ./$(BIN_NAME) $(BIN_PATH)
	rm -rf ./$(BIN_NAME)
	@echo ">>> build $(BIN_NAME) mac platform complete"

.PHONY: linux
linux:
	@echo "build $(BIN_NAME) linux platform >>>"
	rm -rf $(BIN_PATH)
	mkdir -p $(BIN_PATH)
	cp ./.env $(BIN_PATH)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o $(BIN_NAME) ./main.go
	#upx --brute $(BIN_NAME)
	cp ./$(BIN_NAME) $(BIN_PATH)
	rm -rf ./$(BIN_NAME)
	@echo ">>> build $(BIN_NAME) linux platform complete"

.PHONY: windows
windows:
	@echo "build $(BIN_NAME) windows platform >>>"
	rm -rf $(BIN_PATH)
	mkdir -p $(BIN_PATH)
	cp ./.env $(BIN_PATH)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags '-w -s' -o $(BIN_NAME) ./main.go
	#upx --brute $(BIN_NAME)
	cp ./$(BIN_NAME) $(BIN_PATH)
	rm -rf ./$(BIN_NAME)
	@echo ">>> build $(BIN_NAME) windows platform complete"

.PHONY: clean
clean:
	@echo "clean start >>>"
	rm -rf BIN_PATH
	@echo ">>> clean complete"