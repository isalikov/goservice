V := @

OUT_DIR := ./target

MAIN_OUT := $(OUT_DIR)/goservice
MAIN_PKG := ./cmd/goservice

default: build

.PHONY: build
build:
	@echo BUILDING $(MAIN_OUT)
	$(V)go build -o $(MAIN_OUT) $(MAIN_PKG)
	@echo DONE

.PHONY: vendor
vendor:
	$(V)go mod tidy
	$(V)go mod vendor

.PHONY: generate
generate:
	 protoc --go_out=./pkg --go_opt=paths=source_relative \
        --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative \
        api/grpc/goservice.proto

.PHONY: start
start:
	go run ./cmd/goservice/main.go
