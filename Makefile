GO_CMD=go
API_BIN=xingshen_visitor
API_BIN_PATH=./output
API_CONF_PATH=./output
API_LOG_PATH=./output/log

.PHONY: install
install:
	@echo "install huike-api start >>>"
	rm -rf $(API_BIN_PATH)
	mkdir -p $(API_BIN_PATH)
	cp ./app.ini $(API_CONF_PATH)
	cp ./prod.app.ini $(API_CONF_PATH)
	cp ./prod.task.app.ini $(API_CONF_PATH)
	cp ./prod.xs.app.ini $(API_CONF_PATH)
	mkdir -p $(API_LOG_PATH)
	#$(GO_CMD) build -o $(API_BIN) ./main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_CMD) build -ldflags '-w -s' -o $(API_BIN) ./main.go
	#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_CMD) build -o $(API_BIN) ./main.go
	upx --brute $(API_BIN)
	cp ./$(API_BIN) $(API_BIN_PATH)
	rm -rf ./$(API_BIN)
	@echo ">>> install huike-api complete"

.PHONY: clean
clean:
	@echo "clean start >>>"
	rm -fr ./output
	@echo ">>> clean complete"